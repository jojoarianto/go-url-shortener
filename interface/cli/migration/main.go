package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jojoarianto/go-url-shortener/config"
	"github.com/jojoarianto/go-url-shortener/domain/model"
)

func main() {
	conf := config.NewConfig("sqlite3", "url-shortener.sqlite3")
	db, _ := conf.ConnectDB()

	DBMigrate(db)
}

// DBMigrate will create and migrate the tables
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Shorten{})
	log.Println("Schema migration has been procceed")
	return db
}
