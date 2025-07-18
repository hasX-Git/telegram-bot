package main

import (
	"tg-bot/Helpers"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	Helpers.InitServices()
	// Helpers.ConnectToDB()

	Helpers.Bh.Handle(Helpers.Start, th.CommandEqual("start"))
	Helpers.Bh.Handle(Helpers.Help, th.CommandEqual("help"))
	Helpers.Bh.Handle(Helpers.GetFile, th.CommandEqual("getfile"))
	Helpers.Bh.Handle(Helpers.Chat, th.AnyMessage())

	_ = Helpers.Bh.Start()
}
