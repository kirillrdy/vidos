package srt2webvtt

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//TODO looks like this wasnt a very useful helper
func p(err error) {
	if err != nil {
		panic(err)
	}
}

//TODO needs to handle errors
func Convert(srtFilename string, webvttFilename string) {
	//TODO don't read all of the file into memory
	srtData, err := ioutil.ReadFile(srtFilename)

	//TODO handle errors
	p(err)

	//TODO WRONG WRONG WRONG, need to just replace , to . in the timestamp
	replaced := strings.Replace(string(srtData), ",", ".", -1)
	replaced = "WEBVTT FILE\n\n" + replaced

	webvttFile, err := os.Create(webvttFilename)
	//TOOD handle errors
	p(err)
	defer webvttFile.Close()

	io.WriteString(webvttFile, replaced)
}
