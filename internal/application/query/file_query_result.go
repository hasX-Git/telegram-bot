package mapper

import (
	ac "tg-bot/internal/application/common"
)

type FileQueryResult struct {
	Result *ac.FileResult
}

type FileQueryListResult struct {
	Result []*ac.FileResult
}
