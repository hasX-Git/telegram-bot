package main

import (
	H "tg-bot/internal"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	H.InitServices()
	H.ConnectToDB()

	H.Bh.Handle(H.Start, th.CommandEqual("start"))
	H.Bh.Handle(H.Help, th.CommandEqual("help"))
	H.Bh.Handle(H.Info, th.CommandEqual("info"))
	H.Bh.Handle(H.GetFile, th.CommandEqual("getfile"))
	H.Bh.Handle(H.GetAccountInfo, th.CommandEqual("getaccountinfo"))
	H.Bh.Handle(H.GetFileList, th.CommandEqual("getfilelist"))
	H.Bh.Handle(H.GetAccountList, th.CommandEqual("getaccountlist"))
	H.Bh.Handle(H.Message, th.AnyMessage())

	_ = H.Bh.Start()
}
