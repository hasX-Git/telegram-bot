package command

import ac "tg-bot/internal/application/common"

type CreateFileCommand struct {
	Filename string
	Hash     string
	AID      string
	Bytes    []byte
}

type CreateFileResult struct {
	Result *ac.FileResult
}
