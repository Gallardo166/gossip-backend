package controllers

import (
	"encoding/json"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"log"
	"net/http"
)

type Post models.Post

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := initializers.DB.Queryx(
		`SELECT p.id, title, body, image_url, c.name, username
     FROM posts AS p
     JOIN users AS u
     ON p.user_id = u.id
     JOIN categories AS c
     ON p.category_id = c.id`)

	if err != nil {
		log.Fatalf("Error querying data: %s", err)
	}

	var posts []*Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.ImageUrl,
			&post.Category,
			&post.Username,
		)

		if err != nil {
			log.Fatalf("Error scanning data: %s", err)
		}

		posts = append(posts, &post)
	}

	jsonPosts, err := json.Marshal(posts)

	if err != nil {
		log.Fatalf("Error converting to JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonPosts)

	if err != nil {
		log.Fatalf("Error sending JSON: %s", err)
	}
}
