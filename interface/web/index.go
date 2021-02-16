package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	RespondWithJSON(w, http.StatusOK, "Shortener API")
}
