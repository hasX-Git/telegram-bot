package repositories

import "tg-bot/internal/domain/entities"

type FileRepository interface {
	Create(file *entities.File) (*entities.File, error)
	FindByHash(hash string) (*entities.File, error)
	FindAll() ([]*entities.File, error)
	Delete(hash string) error
}
