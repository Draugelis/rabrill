package fetchers

import (
	"rabrill/utils"
	"sync"
)

type Commenter struct {
	ChannelUrl string  `json:"channelUrl"`
	Uploads    []Video `json:"uploads"`
}

type CommenterVideos struct {
	TargetChannel string      `json:"targetChannel"`
	Commenters    []Commenter `json:"commenters"`
}

func addVideos(wg *sync.WaitGroup, caid string, key string, cvids *CommenterVideos) {
	defer wg.Done()
	vids, _ := GetVideoIds(caid, key)
	if vids != nil {
		commenter := Commenter{ChannelUrl: utils.GenChannelUrl(caid)}
		for _, vid := range vids {
			upload, err := GetVideoDetails(vid, key)
			if err != nil {
				continue // ¯\_(ツ)_/¯
			}
			commenter.Uploads = append(commenter.Uploads, upload)
		}
		cvids.Commenters = append(cvids.Commenters, commenter)
	}
}

func FetchCommenterVideos(channel string, key string) (CommenterVideos, error) {
	var commenterVideos CommenterVideos

	cid, err := UrlToId(channel)
	if err != nil {
		return commenterVideos, err
	}
	cids := GetCommentAuthorIds(cid, key)
	cids = utils.Unique(cids)

	commenterVideos.TargetChannel = channel

	var wg sync.WaitGroup
	for _, caid := range cids {
		wg.Add(1)
		go addVideos(&wg, caid, key, &commenterVideos)
	}
	wg.Wait()

	return commenterVideos, nil
}
