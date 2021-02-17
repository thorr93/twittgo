package db

import (
	"context"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertRegister where its register the user into the DB
func InsertRegister(u models.Users) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("twittgo")
	col := bd.Collection("users")
	//I encrypt the password
	u.Password, _ = EncryptPassword(u.Password)
	//I add a single record
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//To obtain the ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//Retorn the object as string
	return ObjID.String(), true, nil
}
