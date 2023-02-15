package utils

import "regexp"

func ValidateUrl(url string) bool {
	re := regexp.MustCompile(`http[s]?:\/\/(www\.)?youtube\.com\/(@|watch\?v=|channel\/)[\w-]+`)
	return re.MatchString(url)
}
