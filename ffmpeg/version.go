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
	result := versionRegex.FindStringSubmatch(string(out))

	//Failed to match
	if len(result) != 2 {
		return ""
	} else {
		return result[1]
	}
}

//Checks that system ffmpeg is of the same version as this lib was built against
//No errors raised just warning is printed
func CheckVersion() {
	if Version() != testedVersion {
		log.Print("WARNING: running against untested version of ffmpeg")
		log.Printf("WARNING: recommended version is %#v", testedVersion)
		log.Printf("WARNING: detected version is %#v", Version())
	}
}
