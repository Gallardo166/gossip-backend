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
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	ImageUrl  *string  `json:"imageUrl,omitempty"`
	Category  string   `json:"category"`
	Username  string   `json:"username"`
	Date      string   `json:"date"`
	LikeCount int      `json:"likeCount"`
	Comments  []string `json:"comments"`
}
