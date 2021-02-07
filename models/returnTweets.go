package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReturnTweets es la estructura usada para los tweet
type ReturnTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userID" json:"userID,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
