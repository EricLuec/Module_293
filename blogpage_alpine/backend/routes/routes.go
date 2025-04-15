package routes

import (
	"blog-app/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	r.HandleFunc("/api/posts/{id}", controllers.GetPost).Methods("GET")
	r.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", controllers.UpdatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("DELETE")
}
