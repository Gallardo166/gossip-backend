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

	"github.com/go-playground/validator/v10"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		helper.WriteError(w, err, http.StatusBadRequest)
		return
	}

	row := initializers.DB.QueryRow(fmt.Sprintf(GetUserByUsername, user.Username))

	var userData models.User

	err = row.Scan(
		&userData.Id,
		&userData.Username,
		&userData.Password,
	)

	//check if username already exists
	if err != nil {
		if err == sql.ErrNoRows {
			helper.WriteJsonError(w, "Wrong username or password", http.StatusBadRequest)
		} else {
			helper.WriteError(w, err, http.StatusInternalServerError)
		}
		//check if password is correct
	} else if !config.Compare(user.Password, userData.Password) {
		helper.WriteJsonError(w, "Wrong username or password", http.StatusBadRequest)
	} else {
		tokenString, err := config.CreateToken(userData.Id)
		if err != nil {
			helper.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		type responseData struct {
			TokenString string `json:"tokenString"`
			Username    string `json:"username"`
			Password    string `json:"password"`
		}

		data := responseData{tokenString, user.Username, user.Password}
		helper.WriteJson(w, data)
	}
}
