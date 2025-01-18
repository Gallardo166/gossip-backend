package controllers

import (
	"database/sql"
	"fmt"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	category := r.URL.Query().Get("category")
	sort := r.URL.Query().Get("sort")
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
	if sort == "time" {
		finalQuery += " ORDER BY date DESC"
	} else {
		finalQuery += " ORDER BY like_count DESC"
	}
	rows, err := initializers.DB.Queryx(finalQuery)

	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}

	var posts []*models.PostPreview
	for rows.Next() {
		var post models.PostPreview
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
			helper.WriteError(w, err, http.StatusInternalServerError)
		}

		posts = append(posts, &post)
	}
	helper.WriteJson(w, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	row := initializers.DB.QueryRow(fmt.Sprintf(GetPostQuery, id))

	var post models.Post
	err := row.Scan(
		&post.Title,
		&post.Body,
		&post.ImageUrl,
		&post.Category,
		&post.Username,
		&post.Date,
		&post.LikeCount,
		&post.CommentCount,
		pq.Array(&post.Comments),
	)

	if err != nil {
		if err == sql.ErrNoRows {
			helper.WriteError(w, fmt.Errorf("no rows match query"), http.StatusBadRequest)
		} else {
			helper.WriteError(w, err, http.StatusInternalServerError)
		}
	}
	helper.WriteJson(w, post)
}
