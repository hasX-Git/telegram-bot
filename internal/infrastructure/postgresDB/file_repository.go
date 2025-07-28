package postgres

import (
	de "tg-bot/internal/domain/entities"
	dr "tg-bot/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormFileRepository struct {
	db *gorm.DB
}

func NewGormFileRepository(db *gorm.DB) dr.FileRepository {
	return &GormFileRepository{db: db}
}

func (cr *GormFileRepository) Create(file *de.File) (*de.File, error) {
	dbFile := toDBfile(file)

	if err := cr.db.Create(dbFile).Error; err != nil {
		return nil, err
	}

	return cr.FindByHash(dbFile.Hash)
}

func (cr *GormFileRepository) FindByHash(hash string) (*de.File, error) {
	var dbFile File
	if err := cr.db.First(&dbFile, hash).Error; err != nil {
		return nil, err
	}

	return fromDBfile(&dbFile), nil
}

func (cr *GormFileRepository) FindAll() ([]*de.File, error) {
	var dbFiles []File

	if err := cr.db.Find(&dbFiles).Error; err != nil {
		return nil, err
	}

	Files := make([]*de.File, len(dbFiles))
	for i, dbFile := range dbFiles {
		Files[i] = fromDBfile(&dbFile)
	}

	return Files, nil
}

func (cr *GormFileRepository) Delete(hash string) error {
	return cr.db.Delete(&File{}, hash).Error
}
