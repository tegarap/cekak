package database

import (
	"github.com/tegarap/cekak/api/models"

	"gorm.io/gorm"
)

type (
	SiteModel struct {
		db *gorm.DB
	}
	SiteModelImpl interface {
		GetAll() ([]models.Site, error)
		Add(url models.Site) (models.Site, error)
		IsExist(shortUrl string) (interface{}, error)
		GetSite(shortUrl string) (models.Site, error)
	}
)

func NewUrlDbModel(db *gorm.DB) *SiteModel {
	//db.AutoMigrate(&SiteModel{})
	return &SiteModel{db: db}
}

func (m *SiteModel) GetAll() ([]models.Site, error) {
	var allUrl []models.Site
	err := m.db.Find(&allUrl).Error
	return allUrl, err
}

func (m *SiteModel) Add(url models.Site) (models.Site, error) {
	err := m.db.Create(&url).Error
	return url, err
}

func (m *SiteModel) IsExist(shortUrl string) (interface{}, error) {
	var url models.Site
	err := m.db.Where("short_url = ?", shortUrl).First(&url).Error
	return url, err
}

func (m *SiteModel) GetSite(shortUrl string) (models.Site, error) {
	var url models.Site
	err := m.db.Where("short_url = ?", shortUrl).First(&url).Error
	return url, err
}