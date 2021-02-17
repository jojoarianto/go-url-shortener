package web

import (
	"net/http"

	"github.com/jojoarianto/go-url-shortener/config"
	"github.com/jojoarianto/go-url-shortener/domain/constant"
	"github.com/jojoarianto/go-url-shortener/infrastructure/sqlite3"
	"github.com/jojoarianto/go-url-shortener/service"
	"github.com/julienschmidt/httprouter"
)

// GetLink method to get redirect link
func GetLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	shortcode := ps.ByName("shortcode")

	// init db
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, constant.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	// call service to validate shorten
	shortenSvc := service.NewShortenService(sqlite3.NewShortenRepo(db))

	// search shortcode
	shorten, err := shortenSvc.GetByShortCode(shortcode)
	if err == nil {
		// todo: add counter

		http.Redirect(w, r, shorten.Url, http.StatusSeeOther)
		return
	}

	if err.Error() == "record not found" {
		RespondWithJSON(w, http.StatusNotFound, response{
			Message:    "The shortcode cannot be found in the system",
			StatusCode: http.StatusNotFound,
		})

		return
	}

	RespondWithError(w, http.StatusInternalServerError, constant.ErrInternalServerError.Error())
}
