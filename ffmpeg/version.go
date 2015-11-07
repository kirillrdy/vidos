package ffmpeg

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

//TODO something better than this, something like min max range
var testedVersions = []string{"2.3.6", "2.6.5", "2.6.4", "2.8.1"}

//Version returns the version of ffmpeg found in path
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
	}
	return result[1]
}

//CheckVersion checks that system ffmpeg is of the same version as this lib was built against
//No errors raised just warning is printed
func CheckVersion() {
	var foundVersion bool
	for _, testedVersion := range testedVersions {
		if Version() == testedVersion {
			foundVersion = true
		}
	}
	if foundVersion == false {
		log.Print("WARNING: running against untested version of ffmpeg")
		log.Printf("WARNING: recommended versions are %#v", strings.Join(testedVersions, ", "))
		log.Printf("WARNING: detected version is %#v", Version())
	}
}
