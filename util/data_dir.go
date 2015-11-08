package util

import (
	"log"
	"os"
	"os/user"
)

//VidosDataDir is where all vidos files will be stored
var VidosDataDir string

//VidosDataDirFor returns a dir where each submodule can store its data
func VidosDataDirFor(suffix string) string {

	dir := VidosDataDir + "/" + suffix

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	return dir
}

func init() {

	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}

	VidosDataDir = user.HomeDir + "/.vidos"
	err = os.MkdirAll(VidosDataDir, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
}
