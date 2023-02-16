package fetchers

import (
	"fmt"
	"rabrill/utils"
)

type Video struct {
	Id           string `json:"id"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	Duration     string `json:"duration"`
	ViewCount    string `json:"views"`
	LikeCount    string `json:"likes"`
	CommentCount string `json:"comments"`
}

type VideoDetails struct {
	Items []struct {
		ContentDetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
		Statistics struct {
			ViewCount    string `json:"viewCount"`
			LikeCount    string `json:"likeCount"`
			CommentCount string `json:"commentCount"`
		} `json:"statistics"`
		Snippet struct {
			Title string `json:"title"`
		}
	} `json:"items"`
}

func GetVideoDetails(vid string, key string) (Video, error) {
	var video Video
	var videoDetails VideoDetails
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=%s&key=%s", vid, key)
	err := utils.MakeRequest(url, &videoDetails)
	if err != nil {
		return video, err
	}

	video.Id = vid
	video.Url = utils.GenVideoUrl(vid)
	video.Title = videoDetails.Items[0].Snippet.Title
	video.Duration = videoDetails.Items[0].ContentDetails.Duration
	video.ViewCount = videoDetails.Items[0].Statistics.ViewCount
	video.LikeCount = videoDetails.Items[0].Statistics.LikeCount
	video.CommentCount = videoDetails.Items[0].Statistics.CommentCount

	return video, nil
}
