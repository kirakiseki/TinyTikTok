package utils

import "time"

/*
redis中的常量
*/
const (
	LoginUserKey      string        = "login:token:"
	PutVideosKey      string        = "videos:put:"
	TotalVideosKey    string        = "videos:total:"
	PersonalVideosKey string        = "videos:Personal:"
	FollowUserKey     string        = "follow:"
	FansUserKey       string        = "fans:"
	VideoCommentKey   string        = "comment:"
	VideoLikeKey      string        = "like:"
	LoginUserTTL      time.Duration = 30 * 10000000000
	CacheNullTTL      time.Duration = 2 * 10000000000
)
