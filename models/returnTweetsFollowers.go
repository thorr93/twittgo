package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReturnTweetsFollowers is the struct that returns the tweets
type ReturnTweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userID,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationID,omitempty"`
	Tweet          struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
