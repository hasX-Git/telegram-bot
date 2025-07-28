package mapper

import (
	ac "tg-bot/internal/application/common"
)

type TransactionQueryResult struct {
	Result *ac.TransactionResult
}

type TransactionQueryListResult struct {
	Result []*ac.TransactionResult
}
