package utils

import "testing"

func TestValidateUrl(t *testing.T) {
	urls := []string{
		"https://www.youtube.com/watch?v=jNQXAC9IVRw",
		"https://youtube.com/watch?v=jNQXAC9IVRw",
		"http://www.youtube.com/watch?v=jNQXAC9IVRw",
		"https://www.youtube.com/@jawed",
		"https://youtube.com/@jawed",
		"http://youtube.com/@jawed",
		"https://www.youtube.com/channel/UC4QobU6STFB0P71PMvOGN5A",
		"https://youtube.com/channel/UC4QobU6STFB0P71PMvOGN5A",
		"http://youtube.com/channel/UC4QobU6STFB0P71PMvOGN5A",
		"http://youtube.com/channel/",
		"http://youtube.com/",
		"youtube.com/channel/",
		"someInvalidString",
	}
	want := []bool{
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
	}

	for i, url := range urls {
		res := ValidateUrl(url)
		if res != want[i] {
			t.Errorf("Failed ValidateUrl(%s). Got: %t. Want: %t", url, res, want[i])
		}
	}
}
