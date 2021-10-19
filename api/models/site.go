package models

type Site struct {
	LongUrl  string `json:"site_url" validate:"required,url"`
	ShortUrl string `json:"short_url" gorm:"unique" validate:"required"`
}
