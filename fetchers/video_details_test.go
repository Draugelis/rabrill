package fetchers

import (
	"encoding/json"
	"testing"

	"github.com/h2non/gock"
)

func TestGetVideoDetails(t *testing.T) {
	defer gock.Off()

	mock := []byte(`
	{
		"kind": "youtube#videoListResponse",
		"etag": "xEY3cMNpcoxXuz2o2E8UEQCRx_U",
		"items": [
			{
				"kind": "youtube#video",
				"etag": "Ck-MdM2DnK0T8t1MUQ7nE9VHVRY",
				"id": "jNQXAC9IVRw",
				"snippet": {
					"publishedAt": "2005-04-24T03:31:52Z",
					"channelId": "UC4QobU6STFB0P71PMvOGN5A",
					"title": "Me at the zoo"
				},
				"contentDetails": {
					"duration": "PT19S"
				},
				"statistics": {
					"viewCount": "256993049",
					"likeCount": "13206712",
					"favoriteCount": "0",
					"commentCount": "11271468"
				}
			}
		],
		"pageInfo": {
			"totalResults": 1,
			"resultsPerPage": 1
		}
	}`)
	var wantedDetails VideoDetails
	json.Unmarshal(mock, &wantedDetails)

	var want Video
	want.Id = "jNQXAC9IVRw"
	want.Url = "https://youtube.com/watch?v=jNQXAC9IVRw"
	want.Title = wantedDetails.Items[0].Snippet.Title
	want.Duration = wantedDetails.Items[0].ContentDetails.Duration
	want.ViewCount = wantedDetails.Items[0].Statistics.ViewCount
	want.LikeCount = wantedDetails.Items[0].Statistics.LikeCount
	want.CommentCount = wantedDetails.Items[0].Statistics.CommentCount

	gock.New("https://youtube.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=jNQXAC9IVRw&key=someKey").
		Reply(200).
		JSON(mock)

	res, err := GetVideoDetails("jNQXAC9IVRw", "someKey")

	if res != want || err != nil {
		t.Errorf("Failed GetVideoDetails('jNQXAC9IVRw', 'someKey'). Got: %v. Want: %v", res, want)
	}

	gock.New("https://youtube.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=errorId&key=someKey").
		Reply(404)

	_, err = GetVideoDetails("errorId", "someKey")
	if err == nil {
		t.Error("Failed GetVideoDetails('errorId', 'someKey'). Got: error=nil. Want: error!=nil")
	}
}
