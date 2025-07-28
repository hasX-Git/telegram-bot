package command

import ac "tg-bot/internal/application/common"

type CreateTransactionCommand struct {
	AID string
	Sum uint32
	TID string
}

type CreateTransactionResult struct {
	Result *ac.TransactionResult
}
