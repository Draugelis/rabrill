package utils

import "testing"

func TestGenChannelUrl(t *testing.T) {
	cid := "UC4QobU6STFB0P71PMvOGN5A"
	want := "https://www.youtube.com/channel/UC4QobU6STFB0P71PMvOGN5A"
	res := GenChannelUrl(cid)
	if res != want {
		t.Errorf("Failed GenChannelUrl(%s). Got: %s. Want: %s", cid, res, want)
	}
}
