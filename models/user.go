package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Email  string             `json:"email" bson:"email"`
	Age    int                `json:"age" bson:"age"`
	Active bool               `json:"active" bson:"active"`
}
