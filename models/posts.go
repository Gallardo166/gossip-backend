package models

type Post struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`
	Body     string  `json:"body"`
	ImageUrl *string `json:"image_url,omitempty"`
	Category string  `json:"category"`
	Username string  `json:"username"`
}
