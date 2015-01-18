package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "dbname=vidos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	Db.AutoMigrate(&Video{})
}
