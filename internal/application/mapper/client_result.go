package mapper

import (
	ac "tg-bot/internal/application/common"
	de "tg-bot/internal/domain/entities"
)

func NewClientResultFromEntity(client *de.Client) *ac.ClientResult {
	if client == nil {
		return nil
	}

	return &ac.ClientResult{
		Firstn: client.Firstn,
		Lastn:  client.Lastn,
		NID:    client.NID,
	}
}
