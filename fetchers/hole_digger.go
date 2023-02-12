package fetchers

import (
	"rabrill/utils"
)

type Upload struct {
	UploadUrl string `json:"videoUrl"`
}

type Commenter struct {
	ChannelUrl string   `json:"channelUrl"`
	Uploads    []Upload `json:"uploads"`
}

type CommenterVideos struct {
	TargetChannel string      `json:"targetChannel"`
	Commenters    []Commenter `json:"commenters"`
}

func FetchCommenterVideos(channel string, key string) CommenterVideos {
	cid := UrlToId(channel)
	cids := GetCommentAuthorIds(cid, key)
	cids = utils.Unique(cids)

	var commenterVideos CommenterVideos
	commenterVideos.TargetChannel = channel

	for _, caid := range cids {
		if caid != cid {
			vids := GetVideoIds(caid, key)
			if vids != nil {
				commenter := Commenter{ChannelUrl: utils.GenChannelUrl(caid)}
				for _, vid := range vids {
					upload := Upload{UploadUrl: utils.GenVideoUrl(vid)}
					commenter.Uploads = append(commenter.Uploads, upload)
				}
				commenterVideos.Commenters = append(commenterVideos.Commenters, commenter)
			}
		}
	}

	return commenterVideos
}
