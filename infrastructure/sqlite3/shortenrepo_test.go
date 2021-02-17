package sqlite3

import (
	"testing"

	"github.com/jojoarianto/go-url-shortener/config"
	"github.com/jojoarianto/go-url-shortener/domain/model"
	"github.com/stretchr/testify/assert"
)

var (
	URIDbConn = "../../url-shortener-test.sqlite3"
	Dialeg    = "sqlite3"
)

func Test_shortenRepo(t *testing.T) {
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		t.Error("error to connect db")
	}
	defer db.Close()

	// test connection
	assert.NoError(t, err)

	// run migration
	db.AutoMigrate(&model.Shorten{})

	// init repo
	repo := NewShortenRepo(db)

	t.Run("success", func(t *testing.T) {
		shorten := model.Shorten{
			Shortcode: "b1234",
			Url:       "http://www.google.com",
		}

		// Test Add
		err := repo.Add(shorten)
		if err != nil {
			t.Error("Error when add")
		}

		assert.NoError(t, err)

		// Test UpdateCount
		err = repo.UpdateRedirectCount(shorten)
		if err != nil {
			t.Error("Error when add")
		}

		assert.NoError(t, err)

		// Test Get
		result, err := repo.GetByShortCode(shorten.Shortcode)
		if err != nil {
			t.Error("Error when add")
		}

		assert.Equal(t, result.Shortcode, shorten.Shortcode)
		assert.Equal(t, result.Url, shorten.Url)

		// assert for update
		assert.Equal(t, result.RedirectCount, (shorten.RedirectCount + 1))
	})

	t.Run("failed", func(t *testing.T) {
		// Test Get when not found
		result, err := repo.GetByShortCode("NOTEXIST")
		if err == nil {
			t.Error("Error when get")
		}

		assert.Equal(t, result.Shortcode, "")
		assert.Equal(t, result.RedirectCount, 0)
	})
}
