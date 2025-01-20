package main

import (
	"gossip-backend/controllers"
	"gossip-backend/initializers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}
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
	//GET posts
	router.Get("/posts", controllers.GetAllPosts)
	router.Get("/post/{id}", controllers.GetPost)
	//POST posts
	router.Post("/post", controllers.PostPost)

	//GET categories
	router.Get("/categories", controllers.GetAllCategories)

	//GET comments
	router.Get("/comments", controllers.GetAllComments)
	//POST comments
	router.Post("/comment", controllers.PostComment)

	//GET users
	router.Get("/user", controllers.GetUser)
	//POST users
	router.Post("/user", controllers.PostUser)

	//authentication
	router.Post("/login", controllers.Login)
	return router
}
