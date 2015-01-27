package srt2webvtt

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func p(err error) {
	if err != nil {
		panic(err)
	}
}

//TODO needs to handle errors
func Convert(srtFilename string, webvttFilename string) {
	//TODO don't read all of the file into memory
	srtData, err := ioutil.ReadFile("public/sub.srt")

	//TODO handle errors
	p(err)

	//TODO WRONG WRONG WRONG, need to just replace , to . in the timestamp
	replaced := strings.Replace(string(srtData), ",", ".", -1)
	replaced = "WEBVTT FILE\n\n" + replaced

	webvttFile, err := os.Create("public/sub.vtt")
	//TOOD handle errors
	p(err)
	defer webvttFile.Close()

	io.WriteString(webvttFile, replaced)
}
