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
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Body       string  `json:"body"`
	ImageUrl   *string `json:"imageUrl,omitempty"`
	CategoryId int     `json:"categoryId"`
	UserId     int     `json:"userId"`
	Date       string  `json:"date"`
}
