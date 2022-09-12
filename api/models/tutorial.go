package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tutorial struct {
	Id          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Published   bool               `json:"published"`
	CreateAt    primitive.DateTime `bson:"Date" json:"createAt"`
	UpdateAt    primitive.DateTime `bason:"Date" json:"updateAt"`
}

func NewTutorial(id primitive.ObjectID, title string, description string, published bool) *Tutorial {
	return &Tutorial{
		Id:          id,
		Title:       title,
		Description: description,
		Published:   published,
	}
}
