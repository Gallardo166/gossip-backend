package main

import (
	"gossip-backend/controllers"
	"gossip-backend/initializers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}

func routers() *chi.Mux {
	//GET posts
	router.Get("/posts", controllers.GetAllPosts)
	router.Get("/post/{id}", controllers.GetPost)
	//POST posts
	router.Post("/post", controllers.PostPost)
	//PUT posts
	router.Put("/post", controllers.UpdatePost)
	//DELETE posts
	router.Delete("/post", controllers.DeletePost)

	//GET categories
	router.Get("/categories", controllers.GetAllCategories)

	//GET comments
	router.Get("/comments", controllers.GetAllComments)
	//POST comments
	router.Post("/comment", controllers.PostComment)
	//PUT comments
	router.Put("/comment", controllers.UpdateComment)
	//DELETE comments
	router.Delete("/comment", controllers.DeleteComment)

	//GET users
	router.Get("/user", controllers.GetUser)
	//POST users
	router.Post("/user", controllers.PostUser)

	//GET likes
	router.Get("/like", controllers.GetLike)
	//POST likes
	router.Post("/like", controllers.PostLike)
	//DELETE likes
	router.Delete("/like", controllers.DeleteLike)

	//authentication
	router.Post("/login", controllers.Login)
	return router
}
