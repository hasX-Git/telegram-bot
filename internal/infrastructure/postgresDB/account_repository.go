package postgres

import (
	de "tg-bot/internal/domain/entities"
	dr "tg-bot/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormAccountRepository struct {
	db *gorm.DB
}

func NewGormAccountRepository(db *gorm.DB) dr.AccountRepository {
	return &GormAccountRepository{db: db}
}

func (cr *GormAccountRepository) Create(account *de.Account) (*de.Account, error) {
	dbAccount := toDBacc(account)

	if err := cr.db.Create(dbAccount).Error; err != nil {
		return nil, err
	}

	return cr.FindByID(dbAccount.NID)
}

func (cr *GormAccountRepository) FindByID(nid string) (*de.Account, error) {
	var dbAccount Account
	if err := cr.db.First(&dbAccount, nid).Error; err != nil {
		return nil, err
	}

	return fromDBacc(&dbAccount), nil
}

func (cr *GormAccountRepository) FindAll() ([]*de.Account, error) {
	var dbAccounts []Account

	if err := cr.db.Find(&dbAccounts).Error; err != nil {
		return nil, err
	}

	Accounts := make([]*de.Account, len(dbAccounts))
	for i, dbAccount := range dbAccounts {
		Accounts[i] = fromDBacc(&dbAccount)
	}

	return Accounts, nil
}

func (cr *GormAccountRepository) Delete(nid string) error {
	return cr.db.Delete(&Account{}, nid).Error
}
