package srt2webvtt

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Convert converts srt subtitle to webvtt format
func Convert(srtFilename string, webvttFilename string) error {
	//TODO don't read all of the file into memory
	srtData, err := ioutil.ReadFile(srtFilename)

	if err != nil {
		return err
	}

	//TODO WRONG WRONG WRONG, need to just replace , to . in the timestamp
	// this also replaces , and . in the text of subtitles
	replaced := strings.Replace(string(srtData), ",", ".", -1)
	replaced = "WEBVTT FILE\n\n" + replaced

	webvttFile, err := os.Create(webvttFilename)
	if err != nil {
		return err
	}

	_, err = io.WriteString(webvttFile, replaced)
	if err != nil {
		return err
	}
	err = webvttFile.Close()
	if err != nil {
		return err
	}
	return nil
}
