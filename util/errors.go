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

// CrashOnErrors panics via the logger if there are any erros
// only use when error is unrecoverable
func CrashOnErrors(err error) {
	if err != nil {
		Errors.Panic(err)
	}
}

func init() {
	Errors = log.New(os.Stderr, "ERROR ", log.LstdFlags)
}
