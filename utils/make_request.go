package utils

import (
	"encoding/json"
	"net/http"
)

func MakeRequest(url string, dst any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(dst)
	return err
}
