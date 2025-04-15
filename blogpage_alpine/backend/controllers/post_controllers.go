package controllers

import (
	"blog-app/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
)

var Client *mongo.Client

func GetPosts(w http.ResponseWriter, r *http.Request) {
	collection := Client.Database("blog").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.M{})
	var posts []models.Post
	cursor.All(ctx, &posts)

	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var post models.Post
	collection := Client.Database("blog").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)

	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	post.Date = time.Now().Format(time.RFC3339)

	collection := Client.Database("blog").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(w).Encode(result)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	update := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"content": post.Content,
		},
	}

	collection := Client.Database("blog").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	w.WriteHeader(http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	collection := Client.Database("blog").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.DeleteOne(ctx, bson.M{"_id": objID})
	w.WriteHeader(http.StatusNoContent)
}
