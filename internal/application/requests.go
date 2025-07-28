package internal

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	u "tg-bot/internal/domain"
	s "tg-bot/internal/infrastructure"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/genai"
	"gorm.io/gorm"
)

type requestType string

const (
	GETFILE  requestType = "getFile"
	GETACC   requestType = "getAcc"
	POSTFILE requestType = "postFile"
	GETSUM   requestType = "getSum"
)

var Requests sync.Map

func getFile(ctx *th.Context, update telego.Update) {

	chatID := update.Message.Chat.ID
	hash := update.Message.Text
	var dbFile u.File

	result := s.DB.First(&dbFile, "hash = ?", hash)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "no such file in database"))
		} else {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when searching for file"))
		}
		return
	}

	filepath := "./files/" + dbFile.Filename

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("error")
		return
	}

	document := tu.Document(
		tu.ID(chatID),
		tu.File(file),
	)

	_, _ = ctx.Bot().SendDocument(ctx, document)
}

func getAcc(ctx *th.Context, update telego.Update) {

	chatID := update.Message.Chat.ID

	AID := update.Message.Text
	var acc u.Account

	result := s.DB.Preload("PersonInfo").Preload("Trs").First(&acc, "aid = ?", AID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("no such file in database")
			return
		} else {
			log.Println("error occured when searching for file")
			return
		}
	}

	var accinfo string

	accinfo +=
		"First Name: " + acc.PersonInfo.Firstn + "\n" +
			"Last Name: " + acc.PersonInfo.Lastn + "\n" +
			"National ID: " + acc.PersonInfo.NID + "\n" +
			"Account ID " + acc.AID + "\n" +
			"Balance: " + strconv.Itoa(int(acc.Balance)) + "\n" +
			"Transactions:\n"

	for _, tr := range acc.Trs {
		accinfo +=
			"\tSum: " + strconv.Itoa(int(tr.Sum)) + "\n" +
				"\tTransaction ID: " + tr.TrID + "\n" +
				"\n"
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), accinfo))

}

func postFile(ctx *th.Context, update telego.Update) {
	chatID := update.Message.Chat.ID

	if update.Message.Document == nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "It has to be a document"))
		return
	}

	var params telego.GetFileParams
	params.FileID = update.Message.Document.FileID

	file, err := s.Bot.GetFile(ctx, &params)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file"))
		return
	}

	urlFile := s.Bot.FileDownloadURL(file.FilePath)
	bytedFile, err := tu.DownloadFile(urlFile)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file's bytes"))
		return
	}

	filepath := "./files/" + update.Message.Document.FileName

	err = os.WriteFile(filepath, bytedFile, 0644)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when downloading file"))
		return
	}

	var newFile u.File
	newFile.Filename = update.Message.Document.FileName
	newFile.Hash = hash(update.Message.Document.FileName)

	result := s.DB.Create(&newFile)

	if result.Error != nil {
		log.Println("Error:", result)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when updating file db"))
		return
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "File uploaded. File hash: "+newFile.Hash))
}

func getSumFile(ctx *th.Context, update telego.Update) {
	chatID := update.Message.Chat.ID

	if update.Message.Document == nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "It has to be a document"))
		return
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Summarizing file..."))

	var params telego.GetFileParams
	params.FileID = update.Message.Document.FileID

	file, err := s.Bot.GetFile(ctx, &params)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file"))
		return
	}

	urlFile := s.Bot.FileDownloadURL(file.FilePath)
	bytedFile, err := tu.DownloadFile(urlFile)
	mimeType := http.DetectContentType(bytedFile)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file's bytes"))
		return
	}

	contents := []*genai.Content{}

	contents = append(contents, genai.Text("Summarize the file")...)
	contents = append(contents, genai.NewContentFromBytes(bytedFile, mimeType, genai.RoleUser))

	summary, err := s.Client.Models.GenerateContent(ctx, s.Aimodel, contents, nil)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when generating summary"))
		return
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), summary.Text()))
}

func RequestExecution(ctx *th.Context, update telego.Update) {
	chatID := update.Message.Chat.ID
	state, _ := Requests.Load(chatID)

	switch state {
	case GETACC:
		getAcc(ctx, update)
	case GETFILE:
		getFile(ctx, update)
	case GETSUM:
		getSumFile(ctx, update)
	case POSTFILE:
		postFile(ctx, update)
	}

}
