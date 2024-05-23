package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"Task,omitempty"`
	Status bool               `json:"Status,omitempty"`
}
