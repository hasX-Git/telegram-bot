package postgres

import (
	de "tg-bot/internal/domain/entities"
	dr "tg-bot/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormClientRepository struct {
	db *gorm.DB
}

func NewGormClientRepository(db *gorm.DB) dr.ClientRepository {
	return &GormClientRepository{db: db}
}

func (cr *GormClientRepository) Create(client *de.Client) (*de.Client, error) {
	dbClient := toDBcl(client)

	if err := cr.db.Create(dbClient).Error; err != nil {
		return nil, err
	}

	return cr.FindByID(dbClient.NID)
}

func (cr *GormClientRepository) FindByID(nid string) (*de.Client, error) {
	var dbClient Client
	if err := cr.db.First(&dbClient, nid).Error; err != nil {
		return nil, err
	}

	return fromDBcl(&dbClient), nil
}

func (cr *GormClientRepository) FindAll() ([]*de.Client, error) {
	var dbClients []Client

	if err := cr.db.Find(&dbClients).Error; err != nil {
		return nil, err
	}

	Clients := make([]*de.Client, len(dbClients))
	for i, dbClient := range dbClients {
		Clients[i] = fromDBcl(&dbClient)
	}

	return Clients, nil
}

func (cr *GormClientRepository) Delete(nid string) error {
	return cr.db.Delete(&Client{}, nid).Error
}
