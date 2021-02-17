package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteTweet - Recibe ID Tweet e ID User
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	//Aqui apunto a la base de datos y la collection
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userID": UserID,
	}
	_, err := col.DeleteOne(ctx, condition)
	return err
}
