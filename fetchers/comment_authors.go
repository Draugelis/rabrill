package fetchers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Comments struct {
	NextPageToken string
	Items         []struct {
		Snippet struct {
			TopLevelComment struct {
				Snippet struct {
					AuthorChannelId struct {
						Value string `json:"value"`
					} `json:"authorChannelId"`
				} `json:"snippet"`
			} `json:"topLevelComment"`
		} `json:"snippet"`
	} `json:"items"`
}

func getComments(url string, key string) Comments {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var comments Comments
	err = decoder.Decode(&comments)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return comments
}

func GetCommentAuthorIds(cid string, key string) []string {
	var comments Comments
	var authorIds []string

	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/commentThreads?maxResults=100&part=snippet&allThreadsRelatedToChannelId=%s&key=%s&pageToken=%s", cid, key, comments.NextPageToken)
		comments = getComments(url, key)
		for _, comment := range comments.Items {
			authorIds = append(authorIds, comment.Snippet.TopLevelComment.Snippet.AuthorChannelId.Value)
		}
		if comments.NextPageToken == "" {
			break
		}
	}

	return authorIds
}
