package repository

import "github.com/jojoarianto/go-url-shortener/domain/model"

type ShortenRepo interface {
	Add(model.Shorten) error
	GetByShortCode(string) (model.Shorten, error)
	UpdateRedirectCount(model.Shorten) error
}
