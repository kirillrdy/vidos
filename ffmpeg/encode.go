package ffmpeg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Encode(inputFilename, outFilename string, progressUpdate func(string)) {
	args := []string{"-y", "-i",
		inputFilename, "-movflags", "faststart", outFilename}

	cmd := exec.Command("ffmpeg", args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Encode()/StderrPipe")
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Println("Encode()/cmd.Start()")
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		splitSlice := strings.Split(text, "=")
		if len(splitSlice) == 2 && splitSlice[0] == "time" {
			progressUpdate(splitSlice[1])
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//TODO on error here, perhaps we can rerun the program with same arguments and
	// Print stderr before it gets consumed by bufio.Scanner
	if err := cmd.Wait(); err != nil {
		log.Println("ffmpeg/Encode()/cmd.Wait")
		log.Printf("try rerunnign: ffmpeg %v\n", strings.Join(args, " "))
		log.Fatal(err)
	}
}
