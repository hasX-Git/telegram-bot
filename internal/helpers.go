package Internal

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/genai"
	"gorm.io/gorm"
)

var fileGet sync.Map

func findFile(ctx *th.Context, update telego.Update) *telego.SendDocumentParams {

	chatID := update.Message.Chat.ID
	hash := update.Message.Text
	var dbFile File

	result := DB.First(&dbFile, "hash = ?", hash)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "no such file in database"))
		} else {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when searching for file"))
		}
		return nil
	}

	filepath := "./files/" + dbFile.Filename

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("error")
		return nil
	}

	document := tu.Document(
		tu.ID(chatID),
		tu.File(file),
	)

	return document
}

var accGet sync.Map

func findAcc(update telego.Update) string {
	AID := update.Message.Text
	var acc Account

	result := DB.Preload("PersonInfo").Preload("Trs").First(&acc, "aid = ?", AID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "no such file in database"
		} else {
			return "error occured when searching for file"
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

	return accinfo

}

var filePost sync.Map

func hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func loadFile(ctx *th.Context, update telego.Update) {
	chatID := update.Message.Chat.ID

	var params telego.GetFileParams
	params.FileID = update.Message.Document.FileID

	file, err := Bot.GetFile(ctx, &params)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file"))
		return
	}

	urlFile := Bot.FileDownloadURL(file.FilePath)
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

	var newFile File
	newFile.Filename = update.Message.Document.FileName
	newFile.Hash = hash(update.Message.Document.FileName)

	result := DB.Create(&newFile)

	if result.Error != nil {
		log.Println("Error:", result)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when updating file db"))
		return
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "File uploaded. File hash: "+newFile.Hash))
}

var fileSum sync.Map

func sumFile(ctx *th.Context, update telego.Update) {
	chatID := update.Message.Chat.ID

	var params telego.GetFileParams
	params.FileID = update.Message.Document.FileID

	file, err := Bot.GetFile(ctx, &params)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file"))
		return
	}

	urlFile := Bot.FileDownloadURL(file.FilePath)
	bytedFile, err := tu.DownloadFile(urlFile)
	mimeType := http.DetectContentType(bytedFile)

	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when getting file's bytes"))
		return
	}

	contents := []*genai.Content{}

	contents = append(contents, genai.Text("Summarize the file")...)
	contents = append(contents, genai.NewContentFromBytes(bytedFile, mimeType, genai.RoleUser))

	summary, err := Client.Models.GenerateContent(ctx, aimodel, contents, nil)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when generating summary"))
		return
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), summary.Text()))
}
