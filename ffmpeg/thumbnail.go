package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
)

//ffmpeg -i input.mp4 -vf  "thumbnail,scale=640:360" -frames:v 1 thumb.png

const ThumbnailWidth = 640
const ThumbnailHeight = 360

//TODO extract 640:360
func Thumbnail(inputFilename, thumFilename string) {

	thumArg := fmt.Sprintf("thumbnail,scale=%v:%v", ThumbnailWidth, ThumbnailHeight)

	cmd := exec.Command("ffmpeg", "-y", "-i",
		inputFilename, "-vf", thumArg, "-frames:v", "1", thumFilename)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
