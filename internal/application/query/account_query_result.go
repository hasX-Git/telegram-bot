package mapper

import (
	ac "tg-bot/internal/application/common"
)

type AccountQueryResult struct {
	Result *ac.AccountResult
}

type AccountQueryListResult struct {
	Result []*ac.AccountResult
}
