package entities

import (
	"errors"
	h "tg-bot/internal/domain/helper"
)

type File struct {
	Filename string
	Hash     string
	AID      string
	Bytes    []byte
}

func NewFile(filename string, aid string, bytes []byte) *File {
	return &File{
		Filename: filename,
		Hash:     h.Hash(filename),
		AID:      aid,
		Bytes:    bytes,
	}
}

func (f *File) validate() error {
	if f.Filename == "" {
		return errors.New("filename must not be empty")
	}
	if f.Hash == "" {
		return errors.New("hash must not be empty")
	}
	if len(f.Bytes) == 0 {
		return errors.New("file must not be empty")
	}
	return nil
}
