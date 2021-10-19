package models

type Site struct {
	LongUrl  string `json:"long_url" validate:"required,url"`
	ShortUrl string `json:"short_url" validate:"required"`
}
