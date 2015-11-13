package ffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

//Duration returns string in format dd:dd:dd.dd
//TODO need better implementation
func Duration(filename string) (string, error) {
	cmd := exec.Command("ffprobe", filename)

	var buffer bytes.Buffer
	cmd.Stderr = &buffer
	err := cmd.Run()
	if err != nil {
		//TODO this is not ideal
		return "", fmt.Errorf("Getting duration of video:  %v", err)
	}

	durationRegex := regexp.MustCompile("Duration: (.*?),")

	result := durationRegex.FindStringSubmatch(buffer.String())

	if len(result) == 2 {
		return result[1], nil
	}

	return "", errors.New("Failed to get version, most likely need to update regex")
}
