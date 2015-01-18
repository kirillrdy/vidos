package ffmpeg

import (
	"io/ioutil"
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
	result := versionRegex.FindStringSubmatch(string(out))

	//Failed to match
	if len(result) != 2 {
		return ""
	} else {
		return result[1]
	}
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

	durationRegex := regexp.MustCompile("Duration: (.*?),")

	output, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err)
	}

	result := durationRegex.FindStringSubmatch(string(output))
	if len(result) == 2 {
		return result[1]
	} else {
		return "ERROR getting duration"
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
