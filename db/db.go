package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Session gorm.DB

func init() {
	var err error
	Session, err = gorm.Open("postgres", "dbname=vidos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	Session.AutoMigrate(&Video{})
}
