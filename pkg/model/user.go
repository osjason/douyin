package model

// User 该对象作为 HTTP 的返回结构体，以及 MongoDB 的存储结构
type User struct {
	Id            int    `json:"id" bson:"id"`
	Name          string `json:"name" bson:"name"`
	FollowCount   int    `json:"follow_Count" bson:"follow_Count"`
	FollowerCount int    `json:"follower_Count" bson:"follower_Count" `
	IsFollow      bool   `json:"is_follow" bson:"is_follow" `
}

type UserSecret struct {
	UserId   int    `bson:"user_id"`
	Password string `bson:"password"`
}

type FavoriteVideoList struct {
	VideoList
}

type FollowList struct {
	UserID   int     `json:"-" bson:"user_id"`
	UserList []*User `json:"user_list" bson:"user_list,inline"`
}

type FollowerList struct {
	UserID   int     `json:"-" bson:"user_id"`
	UserList []*User `json:"user_list" bson:"user_list,inline"`
}
