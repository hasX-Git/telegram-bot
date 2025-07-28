package mapper

import (
	ac "tg-bot/internal/application/common"
)

type ClientQueryResult struct {
	Result *ac.ClientResult
}

type ClientQueryListResult struct {
	Result []*ac.ClientResult
}
