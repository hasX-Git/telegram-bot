package internal

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"

	u "tg-bot/internal/domain"
	s "tg-bot/internal/infrastructure"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/genai"
)

type requestType string

const (
	GETFILE  requestType = "getFile"
	GETACC   requestType = "getAcc"
	POSTFILE requestType = "postFile"
	GETSUM   requestType = "getSum"
)

var Requests sync.Map

func getAcc(aid string) (string, error) {

	acc, err := s.SELECTacc(aid)
	if err != nil {
		return "", errors.New("could not retrieve account")
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

	return accinfo, nil

}

func postFile(ctx *th.Context, update telego.Update) error {
	if update.Message.Document == nil {
		log.Println("It has to be a document")
		return errors.New("it has to be a document")
	}

	var params telego.GetFileParams
	params.FileID = update.Message.Document.FileID

	file, err := s.Bot.GetFile(ctx, &params)
	if err != nil {
		log.Println("error occured when getting file")
		return errors.New("error occured when getting file")
	}

	urlFile := s.Bot.FileDownloadURL(file.FilePath)
	bytedFile, err := tu.DownloadFile(urlFile)
	if err != nil {
		log.Println("error occured when getting file's bytes")
		return errors.New("error occured when getting file's bytes")
	}

	err = s.INSERTfile(bytedFile, update.Message.Document.FileName)
	if err != nil {
		log.Println("error occured when getting file's bytes")
		return errors.New("error occured when getting file's bytes")
	}

	return nil
}

func getSumFile(ctx *th.Context, update telego.Update) string {
	if update.Message.Document == nil {
		return "It has to be a document"
	}

	var params telego.GetFileParams
	params.FileID = update.Message.Document.FileID

	file, err := s.Bot.GetFile(ctx, &params)

	if err != nil {
		log.Println("error occured when getting file")
		return "Could not summarize file"
	}

	urlFile := s.Bot.FileDownloadURL(file.FilePath)
	bytedFile, err := tu.DownloadFile(urlFile)
	mimeType := http.DetectContentType(bytedFile)

	if err != nil {
		log.Println("error occured when getting file's bytes")
		return "Could not summarize file"
	}

	contents := []*genai.Content{}

	contents = append(contents, genai.Text("Summarize the file")...)
	contents = append(contents, genai.NewContentFromBytes(bytedFile, mimeType, genai.RoleUser))

	summary, err := s.Client.Models.GenerateContent(ctx, s.Aimodel, contents, nil)
	if err != nil {
		log.Println("error occured when generating summary")
		return "Could not summarize file"
	}

	return summary.Text()
}

func FileHashList() string {
	var files []u.File

	result := s.DB.Find(&files)

	if result.Error != nil {
		log.Println("Error:", result.Error)
		return "Could not retrieve files"
	}

	var response string

	for _, h := range files {
		response += (h.Hash + ", \n")
	}

	return response

}

func AIDList() string {
	var accs []u.Account

	result := s.DB.Find(&accs)

	if result.Error != nil {
		log.Println("Error:", result.Error)
		return "Could not retrieve accounts"
	}

	var response string

	for _, h := range accs {
		response += (h.AID + ", \n")
	}

	return response
}
