package service

import (
	"testing"

	"github.com/jojoarianto/go-url-shortener/config"
	"github.com/jojoarianto/go-url-shortener/domain/model"
	"github.com/jojoarianto/go-url-shortener/infrastructure/sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	URIDbConn = "../url-shortener-test.sqlite3"
	Dialeg    = "sqlite3"
)

func Test_shortenService_Add(t *testing.T) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	// run migration
	db.AutoMigrate(&model.Shorten{})

	// init repo
	repo := sqlite3.NewShortenRepo(db)

	// Test Add
	shortenSvc := NewShortenService(repo)
	err = shortenSvc.Add(model.Shorten{
		Shortcode: "ABC12",
		Url:       "http://www.google.com",
	})

	assert.NoError(t, err)
}

func Test_shortenService_Validate(t *testing.T) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	// run migration
	db.AutoMigrate(&model.Shorten{})

	// init repo
	repo := sqlite3.NewShortenRepo(db)
	shortenSvc := NewShortenService(repo)

	// Test validates 1
	err = shortenSvc.Validate(model.Shorten{
		Shortcode: "ABC12-----", // INVALID
		Url:       "http://www.google.com",
	})

	assert.Error(t, err)

	// Test validates 2 (Valid Case)
	err = shortenSvc.Validate(model.Shorten{
		Shortcode: "ABC123",
		Url:       "http://www.google.com",
	})

	assert.NoError(t, err)

	// Test validates 3 (Already Exist Shortcode)
	err = shortenSvc.Add(model.Shorten{
		Shortcode: "BBBCCC",
		Url:       "http://www.google.com",
	})
	err = shortenSvc.Validate(model.Shorten{
		Shortcode: "BBBCCC", // Already exist
		Url:       "http://www.google.com",
	})

	assert.Error(t, err)
}

func Test_shortenService_GetByShortCode(t *testing.T) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	// run migration
	db.AutoMigrate(&model.Shorten{})

	// init repo
	repo := sqlite3.NewShortenRepo(db)
	shortenSvc := NewShortenService(repo)

	// Test 1 (Not found)
	result, err := shortenSvc.GetByShortCode("TTTWWW")
	assert.Error(t, err)
	assert.Equal(t, "", result.Shortcode)
	assert.Equal(t, "", result.Url)

	// Test 2 (Not fount)
	err = shortenSvc.Add(model.Shorten{
		Shortcode: "UUUIII",
		Url:       "http://www.google.com/123",
	})
	result, err = shortenSvc.GetByShortCode("UUUIII")
	assert.NoError(t, err)
	assert.Equal(t, "UUUIII", result.Shortcode)
	assert.Equal(t, "http://www.google.com/123", result.Url)
}
