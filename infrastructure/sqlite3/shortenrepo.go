package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/go-url-shortener/domain/model"
	"github.com/jojoarianto/go-url-shortener/domain/repository"
)

type shortenRepo struct {
	Conn *gorm.DB
}

// NewShortenRepo method repo init
func NewShortenRepo(conn *gorm.DB) repository.ShortenRepo {
	return &shortenRepo{Conn: conn}
}

// Add method to add new shorten
func (sr *shortenRepo) Add(shorten model.Shorten) error {
	if err := sr.Conn.Save(&shorten).Error; err != nil {
		return err
	}
	return nil
}

// GetByShortCode method to retrieve a single data shorten by shortcode
func (sr *shortenRepo) GetByShortCode(shortCode string) (model.Shorten, error) {
	data := model.Shorten{}
	if err := sr.Conn.Where("shortcode = ?", shortCode).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Add method to add new shorten
func (sr *shortenRepo) UpdateRedirectCount(shorten model.Shorten) error {
	err := sr.Conn.Model(&shorten).
		Update("redirect_count", (shorten.RedirectCount + 1)).
		Error

	if err != nil {
		return err
	}

	return nil
}
