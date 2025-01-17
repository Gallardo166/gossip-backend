package controllers

import (
	"encoding/json"
	"gossip-backend/config"
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.WriteError(w, err)
	}

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		helper.WriteError(w, err)
	}

	user.Password, err = config.Hash(user.Password)
	if err != nil {
		helper.WriteError(w, err)
	}

	_, err = initializers.DB.NamedExec(PostUserQuery, user)

	if err != nil {
		helper.WriteError(w, err)
	}
}
