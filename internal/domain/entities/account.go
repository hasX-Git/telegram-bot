package entities

import (
	"errors"
	h "tg-bot/internal/domain/helper"
)

type Account struct {
	Balance uint32
	AID     string
	NID     string
}

func NewAccount(nid string) *Account {
	return &Account{
		Balance: 0,
		AID:     "AID" + h.CurrentDateAsID(5),
		NID:     nid,
	}
}

func (a *Account) validate() error {
	if err := h.CheckValidityOfID(a.NID, 12); err != nil {
		return errors.New("invalid Client ID")
	}
	return nil
}
