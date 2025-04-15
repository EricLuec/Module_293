package models

type Post struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string `json:"title,omitempty" bson:"title,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	Date    string `json:"date,omitempty" bson:"date,omitempty"`
}