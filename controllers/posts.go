package controllers

import (
	"database/sql"
	"fmt"
	"gossip-backend/config"
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
			return
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

func PostPost(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	id, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	var post models.InsertPost

	err = r.ParseMultipartForm(2000000)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	post.UserId = id

	data := r.Form
	post.Title = data["title"][0]
	post.Body = data["body"][0]
	post.CategoryId = data["category"][0]
	post.Date = data["date"][0]

	file, header, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if err != http.ErrMissingFile {
		imageUrl, err := config.UploadFile(file, header.Filename)
		if err != nil {
			helper.WriteError(w, err, http.StatusInternalServerError)
			return
		}
		post.ImageUrl = &imageUrl
	}

	_, err = initializers.DB.NamedExec(PostPostQuery, post)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	_, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	var post models.InsertPost

	err = r.ParseMultipartForm(2000000)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	data := r.Form
	post.Id = data["id"][0]
	post.Title = data["title"][0]
	post.Body = data["body"][0]
	post.CategoryId = data["category"][0]
	post.Date = data["date"][0]

	file, header, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if err != http.ErrMissingFile {
		imageUrl, err := config.UploadFile(file, header.Filename)
		if err != nil {
			helper.WriteError(w, err, http.StatusInternalServerError)
			return
		}
		post.ImageUrl = &imageUrl
	}

	_, err = initializers.DB.NamedExec(UpdatePostQuery, post)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}
}
