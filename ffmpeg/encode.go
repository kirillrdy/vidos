package ffmpeg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const ffmpegBinName = "ffmpeg"

//Encode take input file, outputs file and a function which will be called with current ecoded time
func Encode(inputFilename, outFilename string, progressUpdate func(string)) {
	args := []string{"-y", "-i",
		inputFilename,
		"-strict", "-2", //#TODO this is for ffmpeg version that is not compiled with faac, it would be better to detect this in the future
		"-map", "0:v:0",
		"-map", "0:a:0", "-movflags", "faststart", outFilename}

	cmd := exec.Command(ffmpegBinName, args...)

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
		log.Printf("try rerunnign: %v %v\n", ffmpegBinName, strings.Join(args, " "))
		log.Fatal(err)
	}
}
