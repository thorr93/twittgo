package db

import (
	"context"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertTweet saves the Tweet into the DB.
func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("twittgo")
	col := bd.Collection("tweet")

	//We need an exact document like the one that comes in BSON format. Here is the one I am going to save
	register := bson.M{
		"userID":  t.UserID,
		"message": t.Message,
		"fecha":   t.Fecha,
	}

	//We make the insert
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	//If it's okey we return the ID about the tweet created
	objID, _ := result.InsertedID.(primitive.ObjectID) //of the BSON that returns the InsertOne, extracts the ID of the last inserted field and obtains the ObjectID
	return objID.String(), true, nil

}
