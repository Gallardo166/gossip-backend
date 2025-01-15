package controllers

import (
	"fmt"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := initializers.DB.Queryx(GetAllPostsQuery)

	if err != nil {
		helper.WriteError(w, err)
	}

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.ImageUrl,
			&post.Category,
			&post.Username,
			&post.Date,
			&post.LikeCount,
			&post.CommentCount,
		)

		if err != nil {
			helper.WriteError(w, err)
		}

		posts = append(posts, &post)
	}

	helper.WriteJson(w, posts)
}

func GetPostsByTitle(w http.ResponseWriter, r *http.Request) {
	titleParam := chi.URLParam(r, "title")
	fmt.Println(strings.Replace(GetPostsByTitleQuery, "title_query", titleParam, 1))
	rows, err := initializers.DB.Queryx(strings.Replace(GetPostsByTitleQuery, "title_query", titleParam, 1))

	if err != nil {
		helper.WriteError(w, err)
	}

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.ImageUrl,
			&post.Category,
			&post.Username,
			&post.Date,
			&post.LikeCount,
			&post.CommentCount,
		)

		if err != nil {
			helper.WriteError(w, err)
		}

		posts = append(posts, &post)
	}

	helper.WriteJson(w, posts)
}
