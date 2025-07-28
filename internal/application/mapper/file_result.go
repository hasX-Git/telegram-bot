package mapper

import (
	ac "tg-bot/internal/application/common"
	de "tg-bot/internal/domain/entities"
)

func NewFileResultFromEntity(file *de.File) *ac.FileResult {
	if file == nil {
		return nil
	}

	return &ac.FileResult{
		Filename: file.Filename,
		Hash:     file.Hash,
		AID:      file.AID,
		Bytes:    file.Bytes,
	}
}
