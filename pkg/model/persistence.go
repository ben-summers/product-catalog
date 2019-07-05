package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Persistence struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

func NewPersistence() Persistence {
	return Persistence{
		ID: primitive.NewObjectID(),
	}
}
