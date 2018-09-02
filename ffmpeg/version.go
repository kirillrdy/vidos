package ffmpeg

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

//TODO something better than this, something like min max range
var testedVersions = []string{"2.3.6", "2.6.4", "2.6.5", "2.8.1", "2.8.3", "2.8.6", "3.2.4", "3.4", "4.0.2"}

//Version returns the version of ffmpeg found in path
func Version() (string, error) {
	out, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		return "", err
	}

	//TODO maybe this shouldn't be MustCompile
	versionRegex := regexp.MustCompile("version (.*?) Copyright")
	result := versionRegex.FindStringSubmatch(string(out))

	//Failed to match
	if len(result) != 2 {
		//TODO print output for better debugging
		return "", errors.New("Failed to get version from output of ffmpeg -version. report this to maintainer")
	}
	return result[1], nil
}

//CheckVersion checks that system ffmpeg is of the same version as this lib was built against
//No errors raised just warning is printed
func CheckVersion() {

	ffmpegVersion, err := Version()

	if err != nil {
		log.Printf("WARNING: Failed to detect version of ffmpeg: %v", err.Error())
		return
	}

	var foundVersion bool
	for _, testedVersion := range testedVersions {
		if ffmpegVersion == testedVersion {
			foundVersion = true
		}
	}

	if foundVersion == false {
		log.Print("WARNING: running against untested version of ffmpeg")
		log.Printf("WARNING: recommended versions are %#v", strings.Join(testedVersions, ", "))
		log.Printf("WARNING: detected version is %#v", ffmpegVersion)
	}
}
