package mapper

import (
	ac "tg-bot/internal/application/common"
	de "tg-bot/internal/domain/entities"
)

func NewAccountResultFromEntity(account *de.Account) *ac.AccountResult {
	if account == nil {
		return nil
	}

	return &ac.AccountResult{
		Balance: account.Balance,
		AID:     account.AID,
		NID:     account.NID,
	}
}
