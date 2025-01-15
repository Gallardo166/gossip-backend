package controllers

import (
	helper "gossip-backend/helpers"
	"gossip-backend/initializers"
	"gossip-backend/models"
	"net/http"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := initializers.DB.Queryx(GetAllCategoriesQuery)

	if err != nil {
		helper.WriteError(w, err)
	}

	var categories []*models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
		)

		if err != nil {
			helper.WriteError(w, err)
		}

		categories = append(categories, &category)
	}

	helper.WriteJson(w, categories)
}
