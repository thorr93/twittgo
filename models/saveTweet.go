package models

import (
	"time"
)

//SaveTweet is the structure used for tweets
type SaveTweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
