package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}
