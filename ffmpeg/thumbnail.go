package ffmpeg

import (
	"fmt"
	"os/exec"
)

//ffmpeg -i input.mp4 -vf  "thumbnail,scale=640:360" -frames:v 1 thumb.png

const ThumbnailWidth = 640
const ThumbnailHeight = 360

func Thumbnail(inputFilename, thumFilename string) error {

	thumArg := fmt.Sprintf("thumbnail,scale=%v:%v", ThumbnailWidth, ThumbnailHeight)

	args := []string{
		"-y", "-i",
		inputFilename, "-vf", thumArg, "-frames:v", "1", thumFilename,
	}

	cmd := exec.Command("ffmpeg", args...)

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
