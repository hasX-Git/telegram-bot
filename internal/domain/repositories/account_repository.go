package repositories

import "tg-bot/internal/domain/entities"

type AccountRepository interface {
	Create(acc *entities.Account) (*entities.Account, error)
	FindByID(aid string) (*entities.Account, error)
	FindAll() ([]*entities.Account, error)
	Delete(aid string) error
}
