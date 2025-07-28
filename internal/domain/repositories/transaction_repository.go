package repositories

import "tg-bot/internal/domain/entities"

type TransactionRepository interface {
	Create(tr *entities.Transaction) (*entities.Transaction, error)
	FindByID(tid string) (*entities.Transaction, error)
	FindAll() ([]*entities.Transaction, error)
	Update(tr *entities.Transaction) (*entities.Transaction, error)
	Delete(tid string) error
}
