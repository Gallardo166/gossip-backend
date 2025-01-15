package controllers

import (
	"fmt"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	category := r.URL.Query().Get("category")
	var finalQuery string
	switch true {
	case query != "" && category != "":
		finalQuery = fmt.Sprintf(TitleAndCategoryQuery, query, category)
	case query != "" && category == "":
		finalQuery = fmt.Sprintf(TitleQuery, query)
	case query == "" && category != "":
		finalQuery = fmt.Sprintf(CategoryQuery, category)
	default:
		finalQuery = GetAllPostsQuery
	}
	rows, err := initializers.DB.Queryx(finalQuery)

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
