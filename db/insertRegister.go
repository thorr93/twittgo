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
	//Aqui apunto a la base de datos y la collection
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("users")
	//Encripto la contraseña
	u.Password, _ = EncryptPassword(u.Password)
	//Añado un solo registro
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//Para obtener el ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//Retorno el objeto como string
	return ObjID.String(), true, nil
}
