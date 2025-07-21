package main

import (
	Handler "tg-bot/Handlers"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	Handler.InitServices()
	Handler.ConnectToDB()

	Handler.Bh.Handle(Handler.Start, th.CommandEqual("start"))
	Handler.Bh.Handle(Handler.Help, th.CommandEqual("help"))
	Handler.Bh.Handle(Handler.Info, th.CommandEqual("info"))
	Handler.Bh.Handle(Handler.GetFile, th.CommandEqual("getfile"))
	Handler.Bh.Handle(Handler.GetFileList, th.CommandEqual("getfilelist"))
	Handler.Bh.Handle(Handler.Message, th.AnyMessage())

	_ = Handler.Bh.Start()
}
