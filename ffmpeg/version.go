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

//Checks that system ffmpeg is of the same version as this lib was built against
func CheckVersion() {
	if Version() != testedVersion {
		log.Print("WARNING: running against untested version of ffmpeg")
		log.Printf("WARNING: recommended version is %v", testedVersion)
	}
}
