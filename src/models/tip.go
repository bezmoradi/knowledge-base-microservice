package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TipDocument struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Path string             `bson:"path"`
}
