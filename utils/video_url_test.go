package utils

import "testing"

func TestGenVideoUrl(t *testing.T) {
	vid := "jNQXAC9IVRw"
	want := "https://youtube.com/watch?v=jNQXAC9IVRw"
	res := GenVideoUrl(vid)
	if res != want {
		t.Errorf("Failed GenVideoUrl(%s). Got: %s. Want: %s", vid, res, want)
	}
}
