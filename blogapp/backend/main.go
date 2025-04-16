package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}

var collection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("blogdb").Collection("posts")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/posts", postsHandler)
	http.HandleFunc("/api/posts/", postHandler)

	fmt.Println("Server läuft auf http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var posts []Post
		cursor, _ := collection.Find(context.TODO(), bson.M{})
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var post Post
			cursor.Decode(&post)
			posts = append(posts, post)
		}
		json.NewEncoder(w).Encode(posts)
	case http.MethodPost:
		var post Post
		json.NewDecoder(r.Body).Decode(&post)
		collection.InsertOne(context.TODO(), post)
		w.WriteHeader(http.StatusCreated)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/posts/"):]
	objID, _ := primitive.ObjectIDFromHex(id)
	switch r.Method {
	case http.MethodPut:
		var post Post
		json.NewDecoder(r.Body).Decode(&post)
		collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": post})
	case http.MethodDelete:
		collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	}
}
