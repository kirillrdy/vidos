package util

import (
	"log"
	"os"
)

//Errors is where vidos writes errors that are not linked to user actions directly
var Errors *log.Logger

//TODO move this into common util package
//LogError wirtes to Errors if there are any
func LogError(err error) {
	if err != nil {
		Errors.Print(err)
	}
}

func init() {
	Errors = log.New(os.Stderr, "ERROR ", log.LstdFlags)
}
