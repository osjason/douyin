package model

type Comment struct {
	Id          int    `json:"id" bson:"id"`
	VideoId     int    `json:"video_id" bson:"video_id"`
	CommentUser User   `json:"user" bson:"user"`
	Content     string `json:"content" bson:"content"`
	CreateDate  string `json:"create_date" bson:"create_date"`
}

type CommentList struct {
	CommentList []*Comment `json:"comment_list" bson:"comment_list, inline"`
}
