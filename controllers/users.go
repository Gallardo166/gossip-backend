package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gossip-backend/config"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	id, err := config.CheckAuthorized(tokenString)
	if err != nil {
		helper.WriteError(w, err, http.StatusUnauthorized)
	}

	row := initializers.DB.QueryRow(fmt.Sprintf(GetUserById, strconv.Itoa(id)))

	var user models.User
	err = row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			helper.WriteError(w, fmt.Errorf("no rows match query"), http.StatusBadRequest)
		} else {
			helper.WriteError(w, err, http.StatusInternalServerError)
		}
	}
	helper.WriteJson(w, user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.SignupUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		helper.WriteError(w, err, http.StatusBadRequest)
	}

	row := initializers.DB.QueryRow(fmt.Sprintf(GetUserByUsername, user.Username))

	var userData models.User
	errNoUser := row.Scan(
		&userData.Id,
		&userData.Username,
		&userData.Password,
	)

	if errNoUser != nil && errNoUser != sql.ErrNoRows {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	if errNoUser == nil {
		helper.WriteJsonError(w, "This username is taken", http.StatusBadRequest)
		return
	}

	user.Password, err = config.Hash(user.Password)
	if err != nil {
		helper.WriteError(w, err, http.StatusBadRequest)
	}

	_, err = initializers.DB.NamedExec(PostUserQuery, user)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
	}
}
