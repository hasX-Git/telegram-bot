package Internal

import (
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/genai"
)

func Start(ctx *th.Context, update telego.Update) error {

	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("/start"),
			tu.KeyboardButton("/help"),
			tu.KeyboardButton("/info"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("/getfilelist"),
			tu.KeyboardButton("/getFile"),
			tu.KeyboardButton("/getAccountInfo"),
		),
	)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(update.Message.Chat.ID), "Select key").WithReplyMarkup(keyboard))

	return nil
}

func Help(ctx *th.Context, update telego.Update) error {

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(update.Message.Chat.ID), "TODO: List of all commands"))

	return nil
}

func GetFile(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	fileGet.Store(chatID, true)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file hash"))

	return nil
}

func GetAccountInfo(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	accGet.Store(chatID, true)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert account ID"))

	return nil
}

func Message(ctx *th.Context, update telego.Update) error {
	if update.Message != nil {
		chatID := update.Message.Chat.ID

		if isIn, _ := fileGet.Load(chatID); isIn == true {
			_, _ = ctx.Bot().SendDocument(ctx, findFile(ctx, update))
			fileGet.Delete(chatID)
			return nil
		}

		if isIn, _ := accGet.Load(chatID); isIn == true {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), findAcc(update)))
			accGet.Delete(chatID)
			return nil
		}

		if isIn, _ := filePost.Load(chatID); isIn == true {
			loadFile(ctx, update)
			filePost.Delete(chatID)
			return nil
		}

		if isIn, _ := fileSum.Load(chatID); isIn == true {
			sumFile(ctx, update)
			fileSum.Delete(chatID)
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

func GetFileList(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	var files []File

	result := DB.Find(&files)

	if result.Error != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error while getting list"))
		log.Println("Error:", result.Error)
		return result.Error
	}

	var response string

	for _, h := range files {
		response += (h.Hash + ", \n")
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), response))
	return nil
}

func GetAccountList(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	var accs []Account

	result := DB.Find(&accs)

	if result.Error != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error while getting list"))
		log.Println("Error:", result.Error)
		return result.Error
	}

	var response string

	for _, h := range accs {
		response += (h.AID + ", \n")
	}

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), response))
	return nil
}

func Info(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "TODO: Info about bot"))

	return nil
}

func LoadFile(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	filePost.Store(chatID, true)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file and untick \"Compress the image\" option"))

	return nil
}

func FileSummary(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	fileSum.Store(chatID, true)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file"))

	return nil
}
