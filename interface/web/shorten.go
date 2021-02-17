package web

import (
	"encoding/json"
	"net/http"

	"github.com/jojoarianto/go-url-shortener/config"
	"github.com/jojoarianto/go-url-shortener/domain/constant"
	"github.com/jojoarianto/go-url-shortener/domain/model"
	"github.com/jojoarianto/go-url-shortener/infrastructure/sqlite3"
	"github.com/jojoarianto/go-url-shortener/service"
	"github.com/julienschmidt/httprouter"
	"github.com/lucasjones/reggen"
	"gopkg.in/go-playground/validator.v9"
)

type responseShorten struct {
	Shortcode string `json:"shortcode"`
}

// AddShorten method to add shorten
func AddShorten(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	shorten := model.Shorten{}

	// get request payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&shorten); err != nil {
		RespondWithError(w, http.StatusBadRequest, constant.ErrBadParamInput.Error())
		return
	}
	defer r.Body.Close()

	// request validation
	validate := validator.New()
	if err := validate.Struct(shorten); err != nil {
		RespondWithError(w, http.StatusBadRequest, constant.ErrBadParamInput.Error())
		return
	}

	// init db
	conf := config.NewConfig(Dialeg, URIDbConn)
	db, err := conf.ConnectDB()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, constant.ErrInternalServerError.Error())
		return
	}
	defer db.Close()

	// generate shorten code if not exist
	if shorten.Shortcode == "" {
		generateShortcode, err := reggen.Generate("^[0-9a-zA-Z_]{6}$", 10)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, constant.ErrInternalServerError.Error())
			return
		}
		shorten.Shortcode = generateShortcode
	}

	// call service to validate shorten
	shortenSvc := service.NewShortenService(sqlite3.NewShortenRepo(db))
	err = shortenSvc.Validate(shorten)
	if err != nil {
		if err == constant.ErrBadEntity {
			RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
			return
		} else if err == constant.ErrConflict {
			RespondWithError(w, http.StatusConflict, err.Error())
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, constant.ErrInternalServerError.Error())
			return
		}
	}

	// call service to save shorten
	err = shortenSvc.Add(shorten)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, constant.ErrInternalServerError.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, responseShorten{
		Shortcode: shorten.Shortcode,
	})
}
