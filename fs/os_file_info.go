package fs

import (
	"os"
	"path/filepath"
)

func canBeStreamed(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}
	ext := filepath.Ext(file.Name())
	return ext == Mp4
}

//CanBeEncoded for a given os.FileInfo returns if the file can be encoded using ffmpeg
func CanBeEncoded(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}
	ext := filepath.Ext(file.Name())
	if ext == Mp4 || ext == avi || ext == mkv {
		return true
	}
	return false
}

func ChangeExt(name string, newExtension string) string {
	ext := filepath.Ext(name)
	without := name[0 : len(name)-len(ext)]
	return without + newExtension
}
