package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN -It is the object of connection to the DB
var MongoCN = ConnectDB()

//Uso interno
var clientOptions = options.Client().ApplyURI("mongodb+srv://URL_MONGODB")

//ConnectDB - Allos connect to DB
func ConnectDB() *mongo.Client {

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Creada con la DB")
	return client
}

//CheckConnection - Ping to DB
func CheckConnection() bool {
	err := MongoCN.Ping(context.Background(), nil)
	if err != nil {
		return false
	}
	return true
}
