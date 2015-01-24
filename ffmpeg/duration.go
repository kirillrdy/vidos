package ffmpeg

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
)

//TODO need better implementation
// returns string in format dd:dd:dd.dd
func Duration(filename string) string {
	cmd := exec.Command("ffmpeg", "-i", filename)

	var buffer bytes.Buffer
	cmd.Stderr = &buffer
	err := cmd.Run()
	if err != nil {
		log.Printf("Getting duration of video:  %v", err)
	}

	durationRegex := regexp.MustCompile("Duration: (.*?),")

	result := durationRegex.FindStringSubmatch(buffer.String())

	//TODO error handling for this needs to be fixed
	const valueToReturnIfFailed = "00:00:00.00"

	if len(result) == 2 {
		return result[1]
	} else {
		return valueToReturnIfFailed
	}

	return valueToReturnIfFailed
}
