package models

type Comment struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Body       string `json:"body"`
	Date       string `json:"date"`
	ReplyCount int    `json:"replyCount"`
	ParentId   *int   `json:"parentId,omitempty"`
}

type InsertComment struct {
	Id       int    `json:"id" db:"id"`
	Body     string `db:"body" validate:"required,max=800"`
	UserId   int    `db:"userId" validate:"required"`
	PostId   int    `db:"postId" validate:"required"`
	Date     string `db:"date" validate:"required"`
	ParentId int    `db:"parentId"`
}
