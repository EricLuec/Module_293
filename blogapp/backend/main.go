package main

import (
	"blogapp/backend/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", handlers.DeletePost).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
