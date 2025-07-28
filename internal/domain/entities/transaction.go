package entities

import h "tg-bot/internal/domain/helper"

type Transaction struct {
	AID string
	Sum uint32
	TID string
}

func NewTransaction(sum uint32, aid string) *Transaction {
	return &Transaction{
		AID: aid,
		Sum: sum,
		TID: "TID" + h.CurrentDateAsID(5),
	}
}
