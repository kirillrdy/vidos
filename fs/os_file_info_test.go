package fs

import "testing"

func TestChangeExt(t *testing.T) {
	result := ChangeExt("something.avi", ".mp4")
	if result != "something.mp4" {
		t.Errorf("Got %s", result)
	}
}
