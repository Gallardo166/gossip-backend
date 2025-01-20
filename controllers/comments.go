package controllers

import (
	"encoding/json"
	"fmt"
	"gossip-backend/config"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	parentId := r.URL.Query().Get("parentId")
	var finalQuery string
	if parentId == "" {
		finalQuery = GetAllCommentsQuery
	} else {
		finalQuery = fmt.Sprintf(ParentQuery, parentId)
	}
	rows, err := initializers.DB.Queryx(finalQuery)

	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.Id,
			&comment.Username,
			&comment.Body,
			&comment.Date,
			&comment.LikeCount,
			&comment.ReplyCount,
		)

		if err != nil {
			helper.WriteError(w, err, http.StatusInternalServerError)
		}

		comments = append(comments, &comment)
	}

	helper.WriteJson(w, comments)
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	id, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	parentId := r.URL.Query().Get("parentId")
	var finalQuery string
	if parentId == "" {
		finalQuery = PostCommentWithoutParentQuery
	} else {
		finalQuery = PostCommentQuery
	}

	var comment models.InsertComment

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	comment.UserId = id

	_, err = initializers.DB.NamedExec(finalQuery, comment)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	_, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	var comment models.InsertComment

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Print(comment.Body)
	fmt.Print(comment.PostId)

	_, err = initializers.DB.NamedExec(UpdateCommentQuery, comment)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}
}
