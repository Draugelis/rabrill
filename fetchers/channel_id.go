package fetchers

import (
	"log"
	"net/http"
	"os"

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

func UrlToId(url string) string {
	resp, _ := http.Get(url)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var id string
	parseChannelId(doc, &id)

	return id
}
