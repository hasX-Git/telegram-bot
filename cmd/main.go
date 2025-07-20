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
	Handler.Bh.Handle(Handler.GetTestFile, th.CommandEqual("testfile"))
	Handler.Bh.Handle(Handler.GetFile, th.CommandEqual("getfile"))
	Handler.Bh.Handle(Handler.Message, th.AnyMessage())

	_ = Handler.Bh.Start()
}
