package command

import ac "tg-bot/internal/application/common"

type CreateClientCommand struct {
	Firstn string
	Lastn  string
	NID    string
}

type CreateClientResult struct {
	Result *ac.ClientResult
}
