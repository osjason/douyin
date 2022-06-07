package model

import (
	_ "gorm.io/gorm"
)

type Video struct {
	Id            int    `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" `
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	Title         string `json:"title"`
}
type VideoList struct {
	VideoList []Video `json:"video_list"`
}

type FollowerVideoList struct {
	VideoId int
	UserId  int
}
