package Handler

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"gorm.io/gorm"
)

var fileHash sync.Map

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
