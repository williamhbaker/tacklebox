package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// HookDocument is a BSON representation of a received webhook for storing in mongodb.
type HookDocument struct {
	ID      *primitive.ObjectID `bson:"_id"`
	Content string              `bson:"content"`
}
