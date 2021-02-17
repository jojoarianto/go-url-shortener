package service

import "github.com/jojoarianto/go-url-shortener/domain/model"

// ShortenService contract
type ShortenService interface {
	Add(model.Shorten) error
	Validate(model.Shorten) error
	GetByShortCode(string) (model.Shorten, error)
}
