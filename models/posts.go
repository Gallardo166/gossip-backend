package models

type PostPreview struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Body         string  `json:"body"`
	ImageUrl     *string `json:"imageUrl,omitempty"`
	Category     string  `json:"category"`
	Username     string  `json:"username"`
	Date         string  `json:"date"`
	LikeCount    int     `json:"likeCount"`
	CommentCount int     `json:"commentCount"`
}

type Post struct {
	Title        string   `json:"title"`
	Body         string   `json:"body"`
	ImageUrl     *string  `json:"imageUrl,omitempty"`
	Category     string   `json:"category"`
	Username     string   `json:"username"`
	Date         string   `json:"date"`
	LikeCount    int      `json:"likeCount"`
	CommentCount int      `json:"commentCount"`
	Comments     []string `json:"comments"`
}

type InsertPost struct {
	Id         int     `json:"id" db:"id"`
	Title      string  `json:"title" db:"title" validate:"required,max=200"`
	Body       string  `json:"body" db:"body"`
	ImageUrl   *string `json:"imageUrl,omitempty" db:"imageUrl"`
	CategoryId int     `json:"categoryId" db:"categoryId"`
	UserId     int     `json:"userId" db:"userId"`
	Date       string  `json:"date" db:"date"`
}
