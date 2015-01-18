package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// DB is not a good name, since sql.DB represents a connection manager
var Session gorm.DB

func init() {
	var err error
	Session, err = gorm.Open("postgres", "dbname=vidos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	Session.AutoMigrate(&Video{})
}
