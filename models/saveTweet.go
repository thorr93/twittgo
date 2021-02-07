package models

import (
	"time"
)

//SaveTweet es la estructura usada para los tweet
type SaveTweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
