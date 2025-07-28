package repositories

import "tg-bot/internal/domain/entities"

type ClientRepository interface {
	Create(cl *entities.Client) (*entities.Client, error)
	FindByID(nid string) (*entities.Client, error)
	FindAll() ([]*entities.Client, error)
	Delete(nid string) error
}
