package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	URIDbConn = "url-shortener.sqlite3"
	Dialeg    = "sqlite3"
)

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", Index)
	r.POST("/shorten", AddShorten)
	r.GET("/:shortcode", GetLink)
	r.GET("/:shortcode/stats", GetStats)

	return r
}
