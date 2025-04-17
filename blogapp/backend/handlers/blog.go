package handlers

import (
	"blogapp/backend/models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var collection *mongo.Collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("blogdb").Collection("posts")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	post.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": post})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
