package postgres

import (
	de "tg-bot/internal/domain/entities"
	dr "tg-bot/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewGormTransactionRepository(db *gorm.DB) dr.TransactionRepository {
	return &GormTransactionRepository{db: db}
}

func (cr *GormTransactionRepository) Create(transaction *de.Transaction) (*de.Transaction, error) {
	dbTransaction := toDBtr(transaction)

	if err := cr.db.Create(dbTransaction).Error; err != nil {
		return nil, err
	}

	return cr.FindByID(dbTransaction.TID)
}

func (cr *GormTransactionRepository) FindByID(tid string) (*de.Transaction, error) {
	var dbTransaction Transaction
	if err := cr.db.First(&dbTransaction, tid).Error; err != nil {
		return nil, err
	}

	return fromDBtr(&dbTransaction), nil
}

func (cr *GormTransactionRepository) FindAll() ([]*de.Transaction, error) {
	var dbTransactions []Transaction

	if err := cr.db.Find(&dbTransactions).Error; err != nil {
		return nil, err
	}

	Transactions := make([]*de.Transaction, len(dbTransactions))
	for i, dbTransaction := range dbTransactions {
		Transactions[i] = fromDBtr(&dbTransaction)
	}

	return Transactions, nil
}

func (cr *GormTransactionRepository) Update(transaction *de.Transaction) (*de.Transaction, error) {
	dbTransaction := toDBtr(transaction)
	err := cr.db.Model(&Transaction{}).Where("tid = ?", dbTransaction.TID).Updates(dbTransaction).Error
	if err != nil {
		return nil, err
	}

	return cr.FindByID(dbTransaction.TID)
}

func (cr *GormTransactionRepository) Delete(nid string) error {
	return cr.db.Delete(&Transaction{}, nid).Error
}
