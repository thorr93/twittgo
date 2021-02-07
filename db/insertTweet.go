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
	//Aqui apunto a la base de datos y la collection
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("tweet")

	//Necesitamos un documento exacto como el que viene en formato BSON. Aqui se guarda el que voy a guardar
	register := bson.M{
		"userID":  t.UserID,
		"message": t.Message,
		"fecha":   t.Fecha,
	}

	//hacemos el insert
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	//Si es okey devolvemos un ID del tweet creado
	objID, _ := result.InsertedID.(primitive.ObjectID) //Del BSON que devuelve el InsertOne, extrae el ID del ultimo campo insertado y obtiene el ObjectID
	return objID.String(), true, nil

}
