package fetchers

import (
	"fmt"
	"rabrill/utils"
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

func GetVideoIds(channelId string, key string) ([]string, error) {
	pid, err := getChannelPlaylist(channelId, key)
	if err != nil {
		return nil, err
	}
	vids, err := getUploads(pid, key)
	if err != nil {
		return nil, err
	}
	return vids, nil
}

func getChannelPlaylist(cid string, key string) (string, error) {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/channels?part=contentDetails&id=%s&key=%s", cid, key)
	var channel Channel
	err := utils.MakeRequest(url, &channel)
	if err != nil {
		return "", err
	}
	return channel.Items[0].ContentDetails.RelatedPlaylist.Uploads, nil
}

func getUploads(pid string, key string) ([]string, error) {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/playlistItems?maxResults=50&part=contentDetails&playlistId=%s&key=%s", pid, key)
	var playlist Playlist
	err := utils.MakeRequest(url, &playlist)
	if err != nil {
		return nil, err
	}

	var videoIds []string
	for _, video := range playlist.Items {
		videoIds = append(videoIds, video.ContentDetails.VideoId)
	}

	return videoIds, nil
}
