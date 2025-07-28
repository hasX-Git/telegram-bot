package botapi

import (
	"log"

	r "tg-bot/internal/application"
	s "tg-bot/internal/infrastructure"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/genai"
)

func Start(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID

	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("/start"),
			tu.KeyboardButton("/help"),
			tu.KeyboardButton("/info"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("/getfilelist"),
			tu.KeyboardButton("/getFile"),
			tu.KeyboardButton("/loadfile"),
		),
	)

	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Select key").WithReplyMarkup(keyboard))

	return nil
}

func Help(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "TODO: List of all commands"))
	return nil
}

func GetFileList(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), r.FileHashList()))
	return nil
}

func GetAccountList(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), r.AIDList()))
	return nil
}

func Info(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "TODO: Info about bot"))
	return nil
}

func GetFile(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	r.Requests.Store(chatID, r.GETFILE)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file hash"))
	return nil
}

func GetAccountInfo(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	r.Requests.Store(chatID, r.GETACC)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert account ID"))
	return nil
}

func LoadFile(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	r.Requests.Store(chatID, r.POSTFILE)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file and untick \"Compress the image\" option"))
	return nil
}

func FileSummary(ctx *th.Context, update telego.Update) error {
	chatID := update.Message.Chat.ID
	r.Requests.Store(chatID, r.GETSUM)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "Insert file"))
	return nil
}

func Response(ctx *th.Context, update telego.Update) error {
	if update.Message != nil {
		chatID := update.Message.Chat.ID

		if _, exists := r.Requests.Load(chatID); exists {
			//r.RequestExecution(update)
		} else {
			text := "\"" + update.Message.Text + "\"\n"

			response, err := s.Client.Models.GenerateContent(ctx, s.Aimodel, genai.Text(text+s.Prompt), nil)
			if err != nil {
				log.Println(err)
				return err
			}

			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), response.Text()))
		}

		r.Requests.Delete(chatID)
	}

	return nil
}
