package controllers

import (
	"database/sql"
	"fmt"
	"gossip-backend/config"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"net/http"
)

func GetLike(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	id, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}
	postId := r.URL.Query().Get("postId")
	fmt.Printf(GetLikeQuery, id, postId)
	row := initializers.DB.QueryRow(fmt.Sprintf(GetLikeQuery, id, postId))

	var like struct {
		userId int
		postId int
	}

	err = row.Scan(
		&like.userId,
		&like.postId,
	)

	type responseData struct {
		IsLiked bool `json:"isLiked"`
	}

	//check if user has liked the post
	if err != nil {
		if err == sql.ErrNoRows {
			data := responseData{false}
			helper.WriteJson(w, data)
			return
		} else {
			helper.WriteError(w, err, http.StatusInternalServerError)
			return
		}
	}
	data := responseData{true}
	helper.WriteJson(w, data)
}

func PostLike(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	id, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	postId := r.URL.Query().Get("postId")

	_, err = initializers.DB.Exec(fmt.Sprintf(PostLikeQuery, id, postId))
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}
}

func DeleteLike(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	id, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
		return
	}

	postId := r.URL.Query().Get("postId")

	_, err = initializers.DB.Exec(fmt.Sprintf(DeleteLikeQuery, id, postId))
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}
}
