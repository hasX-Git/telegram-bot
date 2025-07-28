package command

import ac "tg-bot/internal/application/common"

type CreateAccountCommand struct {
	Balance uint32
	AID     string
	NID     string
}

type CreateAccountResult struct {
	Result *ac.AccountResult
}
