package database

import (
	"github.com/tegarap/cekak/api/models"
	"gorm.io/gorm"
)

type UrlDbModel struct {
	db *gorm.DB
}

func NewUrlDbModel(db *gorm.DB) *UrlDbModel {
	db.AutoMigrate(&UrlDbModel{})
	return &UrlDbModel{
		db: db,
	}
}

func (m *UrlDbModel) GetAll() ([]models.Url, error) {
	var allUrl []models.Url
	err := m.db.Find(&allUrl).Error
	return allUrl, err
}

func (m *UrlDbModel) Add(url models.Url) (models.Url, error) {
	err := m.db.Save(&url).Error
	return url, err
}

