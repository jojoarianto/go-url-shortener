package service

import (
	"regexp"

	"github.com/jojoarianto/go-url-shortener/domain/constant"
	"github.com/jojoarianto/go-url-shortener/domain/model"
	"github.com/jojoarianto/go-url-shortener/domain/repository"
)

type shortenService struct {
	shortenRepo repository.ShortenRepo
}

// NewShortenService method service init
func NewShortenService(shortenRepo repository.ShortenRepo) *shortenService {
	return &shortenService{
		shortenRepo: shortenRepo,
	}
}

// Add service to create new shorten
func (ss *shortenService) Add(shorten model.Shorten) error {
	err := ss.shortenRepo.Add(shorten)
	if err != nil {
		return err
	}

	return nil
}

// Validate service to create new shorten
func (ss *shortenService) Validate(shorten model.Shorten) error {

	if shorten.Shortcode == "" {
		return nil
	}

	var validShortCode = regexp.MustCompile(`^[0-9a-zA-Z_]{6}$`)

	isValidShortCode := validShortCode.MatchString(shorten.Shortcode)
	if !isValidShortCode {
		return constant.ErrBadEntity
	}

	shortenExist, err := ss.shortenRepo.GetByShortCode(shorten.Shortcode)
	if (err != nil) && (err.Error() != "record not found") {
		return constant.ErrInternalServerError
	}

	if (model.Shorten{}) != shortenExist {
		return constant.ErrConflict
	}

	return nil
}

// GetByShortCode service to create new shorten
func (ss *shortenService) GetByShortCode(shortcode string) (shorten model.Shorten, err error) {
	shorten, err = ss.shortenRepo.GetByShortCode(shortcode)
	if err != nil {
		return
	}

	err = ss.shortenRepo.UpdateRedirectCount(shorten)
	return
}
