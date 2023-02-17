package fetchers

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func parseChannelId(n *html.Node, id *string) {
	if n.Type == html.ElementNode && n.Data == "meta" {
		for i, element := range n.Attr {
			if element.Val == "channelId" {
				*id = n.Attr[i+1].Val
				return
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseChannelId(c, id)
	}
}

func UrlToId(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode > 299 {
		return "", fmt.Errorf("failed to get channel ID; %v", resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var id string
	parseChannelId(doc, &id)

	return id, nil
}
