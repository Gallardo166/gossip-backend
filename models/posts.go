package models

type Post struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Body         string  `json:"body"`
	ImageUrl     *string `json:"image_url,omitempty"`
	Category     string  `json:"category"`
	Username     string  `json:"username"`
	Date         string  `json:"date"`
	LikeCount    int     `json:"like_count"`
	CommentCount int     `json:"comment_count"`
}
