package db

import (
// "log"

// "github.com/jinzhu/gorm"

//sql packages requires pq to register the driver
//_ "github.com/lib/pq"
)

// DB is not a good name, since sql.DB represents a connection manager
// var Postgres gorm.DB

// func init() {
// 	var err error
// 	Postgres, err = gorm.Open("postgres", "dbname=vidos sslmode=disable")
// 	if err != nil {
// 		log.Println("db/init()")
// 		log.Fatal(err)
// 	}

// 	Postgres.AutoMigrate(&Video{}, &Subtitle{})
// }
