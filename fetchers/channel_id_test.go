package fetchers

import (
	"strings"
	"testing"

	"github.com/h2non/gock"
)

func TestUrlToId(t *testing.T) {
	defer gock.Off()

	successBody := strings.NewReader(`<html>
		<head>
			<title>Some content</title>
		</head>
		<body>
			<div>
            	<meta itemprop="channelId" content="UC4QobU6STFB0P71PMvOGN5A">
        	</div>
    	</body>
	</html>`)

	failBody := strings.NewReader("")

	gock.New("https://youtube.com/@successRequest").
		Reply(200).
		Body(successBody)

	res, _ := UrlToId("https://youtube.com/@successRequest")
	want := "UC4QobU6STFB0P71PMvOGN5A"
	if res != want {
		t.Errorf("Failed UrlToId('https://youtube.com/@successRequest'). Get: %s. Want: %s", res, want)
	}

	gock.New("https://youtube.com/@404Request").
		Reply(404)

	_, err := UrlToId("https://youtube.com/@404Request")
	if err == nil {
		t.Errorf("Failed UrlToId('https://youtube.com/@404Request').")
	}

	gock.New("https://youtube.com/@failedRequest").
		Reply(200).
		Body(failBody)

	_, err = UrlToId("https://youtube.com/@failedRequest")
	if err == nil {
		t.Errorf("Failed UrlToId('https://youtube.com/@failedRequest').")
	}

}
