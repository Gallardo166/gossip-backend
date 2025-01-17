package controllers

import (
	"fmt"
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
		helper.WriteError(w, err)
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
			helper.WriteError(w, err)
		}

		comments = append(comments, &comment)
	}

	helper.WriteJson(w, comments)
}
