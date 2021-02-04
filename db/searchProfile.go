package db

import (
	"context"
	"fmt"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchUser - Check a User into the DB
func SearchUser(ID string) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittgo")
	col := db.Collection("users")

	var profile models.Users
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not find" + err.Error())
		return profile, err
	}
	return profile, nil

}
