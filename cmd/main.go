package main

import (
	s "tg-bot/internal/infrastructure"
	r "tg-bot/internal/interface"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	s.Init()

	s.Bh.Handle(r.Start, th.CommandEqual("start"))
	s.Bh.Handle(r.Help, th.CommandEqual("help"))
	s.Bh.Handle(r.Info, th.CommandEqual("info"))
	s.Bh.Handle(r.GetFile, th.CommandEqual("getfile"))
	s.Bh.Handle(r.LoadFile, th.CommandEqual("loadfile"))
	s.Bh.Handle(r.FileSummary, th.CommandEqual("sumfile"))
	s.Bh.Handle(r.GetAccountInfo, th.CommandEqual("getaccountinfo"))
	s.Bh.Handle(r.GetFileList, th.CommandEqual("getfilelist"))
	s.Bh.Handle(r.GetAccountList, th.CommandEqual("getaccountlist"))
	s.Bh.Handle(r.Response, th.AnyMessage())

	_ = s.Bh.Start()
}
