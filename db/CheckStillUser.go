package db

import (
	"context"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckStillUser recibe un email i comprueba si esta en la base de datos
func CheckStillUser(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	//Aqui apunto a la base de datos y la collection
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("users")

	//Condicion en formato bson
	condition := bson.M{"email": email}
	var result models.Users
	//Me devuelve un registro. El decode me convierte los resultados y lo coloco como puntero en resultado
	err := col.FindOne(ctx, condition).Decode(&result)
	//En el tercer parametro, se trae el objectID i aqui se convierte a string
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
