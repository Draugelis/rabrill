package utils

import "fmt"

func GenVideoUrl(vid string) string {
	return fmt.Sprintf("https://youtube.com/watch?v=%s", vid)
}
