package db

import (
	"context"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckStillUser receive an email and check if it is in the database
func CheckStillUser(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("twittgo")
	col := bd.Collection("users")

	//Condition in bson format
	condition := bson.M{"email": email}
	var result models.Users
	//It returns a record. The decode converts the results and I place it as a pointer in result.
	err := col.FindOne(ctx, condition).Decode(&result)
	//In the third parameter, we get the objectID and here it is converted to string
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
