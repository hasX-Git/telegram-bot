package mapper

import (
	ac "tg-bot/internal/application/common"
	de "tg-bot/internal/domain/entities"
)

func NewTransactionResultFromEntity(transaction *de.Transaction) *ac.TransactionResult {
	if transaction == nil {
		return nil
	}

	return &ac.TransactionResult{
		AID: transaction.AID,
		Sum: transaction.Sum,
		TID: transaction.TID,
	}
}
