package main

import (
	"gossip-backend/controllers"
	"gossip-backend/initializers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func init() {
	initializers.ConnectDB()
}

var router *chi.Mux

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	routers()
	http.ListenAndServe(":3000", router)
}

func routers() *chi.Mux {
	router.Get("/posts", controllers.GetAllPosts)
	return router
}
