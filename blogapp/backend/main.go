package main

import (
	"blogapp/backend/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Routen für Posts
	r.HandleFunc("/api/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", handlers.DeletePost).Methods("DELETE")

	// Routen für Kommentare
	r.HandleFunc("/api/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/api/comments/{id}", handlers.UpdateComment).Methods("PUT")
	r.HandleFunc("/api/comments/{id}", handlers.DeleteComment).Methods("DELETE")
	r.HandleFunc("/api/comments/{id}", handlers.GetComments).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
