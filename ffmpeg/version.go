package ffmpeg

import (
	"log"
	"os/exec"
	"regexp"
)

const testedVersion = "2.3.6"

func Version() string {
	out, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		log.Fatal(err)
	}

	versionRegex := regexp.MustCompile("version (.*?) Copyright")
	val := versionRegex.FindStringSubmatch(string(out))[1]
	return val
}

//TODO need better implementation
func Duration(filename string) string {
	cmd := exec.Command("ffmpeg", "-i", filename)

	//TODO does it need to be closed ?
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Panic(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Panic(err)
	}

	//	err = cmd.Wait()
	//	if err != nil {
	//		log.Print(err)
	//	}

	buf := make([]byte, 1024)

	//TODO find a better way
	n := 1
	for n != 0 {
		//TODO handle errors
		n, _ = stderr.Read(buf)
		durationRegex := regexp.MustCompile("Duration: (.*?),")
		result := durationRegex.FindStringSubmatch(string(buf))
		if len(result) == 2 {
			return result[1]
		}
	}

	//TODO not correct
	return ""
}

//Checks that system ffmpeg is of the same version as this lib was built against
func CheckVersion() {
	if Version() != testedVersion {
		log.Print("WARNING: running against untested version of ffmpeg")
		log.Printf("WARNING: recommended version is %v", testedVersion)
	}
}
