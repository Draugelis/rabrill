package fetchers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Channel struct {
	Items []struct {
		ContentDetails struct {
			RelatedPlaylist struct {
				Uploads string `json:"uploads"`
			} `json:"relatedPlaylists"`
		} `json:"contentDetails"`
	} `json:"items"`
}

type Playlist struct {
	Items []struct {
		ContentDetails struct {
			VideoId string `json:"videoId"`
		} `json:"contentDetails"`
	} `json:"items"`
}

func GetVideoIds(channelId string, key string) []string {
	return getUploads(getChannelPlaylist(channelId, key), key)
}

func getChannelPlaylist(cid string, key string) string {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/channels?part=contentDetails&id=%s&key=%s", cid, key)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var channel Channel
	err = decoder.Decode(&channel)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return channel.Items[0].ContentDetails.RelatedPlaylist.Uploads
}

func getUploads(pid string, key string) []string {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/playlistItems?maxResults=50&part=contentDetails&playlistId=%s&key=%s", pid, key)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var playlist Playlist
	err = decoder.Decode(&playlist)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var videoIds []string
	for _, video := range playlist.Items {
		videoIds = append(videoIds, video.ContentDetails.VideoId)
	}

	return videoIds
}
