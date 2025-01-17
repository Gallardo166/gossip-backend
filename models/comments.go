package models

type Comment struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Body       string `json:"body"`
	Date       string `json:"date"`
	LikeCount  int    `json:"likeCount"`
	ReplyCount int    `json:"replyCount"`
	ParentId   *int   `json:"parentId,omitempty"`
}
