package entities

import (
	"errors"
	h "tg-bot/internal/domain/helper"
)

type Client struct {
	Firstn string
	Lastn  string
	NID    string
}

func NewClient(firstn string, lastn string, nid string) *Client {
	return &Client{
		Firstn: firstn,
		Lastn:  lastn,
		NID:    nid,
	}
}

func (c *Client) validate() error {
	if c.Firstn == "" || c.Lastn == "" {
		return errors.New("name must not be empty")
	}
	if c.Firstn == "" {
		return errors.New("name must not be empty")
	}
	if err := h.CheckValidityOfID(c.NID, 12); err != nil {
		return errors.New("invalid ID")
	}
	return nil
}
