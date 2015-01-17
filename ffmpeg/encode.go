package ffmpeg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Encode(inputFilename, outFilename string) {
	cmd := exec.Command("ffmpeg", "-i",
		inputFilename, "-movflags", "faststart", outFilename)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		//log.Print(text)
		splitSlice := strings.Split(text, "=")
		if len(splitSlice) == 2 && splitSlice[0] == "time" {
			log.Printf("%v: %v", inputFilename, splitSlice[1])
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
