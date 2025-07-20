package Handler

import (
	"log"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/genai"
)

const prompt = "If the user seems to ask for some command, offer \"/help\" command ONLY. Otherwise, just respond without offering the command\n"
const aimodel = "gemini-2.5-flash"

func Start(ctx *th.Context, update telego.Update) error {

	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("/start"),
			tu.KeyboardButton("/help"),
			tu.KeyboardButton("/info"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("/insertFile"),
			tu.KeyboardButton("/getFile"),
			tu.KeyboardButton("/getAccountInfo"),
		),
	)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(update.Message.Chat.ID), "Select key").WithReplyMarkup(keyboard))

	return nil
}

func Help(ctx *th.Context, update telego.Update) error {

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(update.Message.Chat.ID), "List of all commands"))

	return nil
}

func GetTestFile(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID

	file, err := os.Open("files/test.txt")
	if err != nil {
		log.Println("error")
		return err
	}

	document := tu.Document(
		tu.ID(chatID),
		tu.File(file),
	)

	_, _ = ctx.Bot().SendDocument(ctx, document)

	return nil
}

func GetFile(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	fileHash.Store(chatID, true)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file hash"))

	return nil
}

func Message(ctx *th.Context, update telego.Update) error {
	if update.Message != nil {
		chatID := update.Message.Chat.ID

		if isIn, _ := fileHash.Load(chatID); isIn == true {
			_, _ = ctx.Bot().SendDocument(ctx, findFile(ctx, update))
			fileHash.Delete(chatID)
			return nil
		}

		text := "\"" + update.Message.Text + "\"\n"

		response, err := Client.Models.GenerateContent(ctx, aimodel, genai.Text(text+prompt), nil)
		if err != nil {
			log.Println(err)
			return err
		}

		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), response.Text()))
	}
	return nil
}
