package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct{
	ID primitive.objectID 'json:"_ID,omitempy" bson: "_id,omitempy"'
	Task string			  'json:"Task,omitempy"'
	Status bool			  'json:"Status,omitempy"'
}