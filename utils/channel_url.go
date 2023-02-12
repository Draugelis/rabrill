package utils

import "fmt"

func GenChannelUrl(cid string) string {
	return fmt.Sprintf("https://www.youtube.com/channel/%s", cid)
}
