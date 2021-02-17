package db

import (
	"context"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModifyRegister allows modify the register of the user
func ModifyRegister(u models.Users, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("twittgo")
	col := bd.Collection("users")

	//make allos create maps or slices
	register := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		register["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		register["apellidos"] = u.Apellidos
	}
	if len(u.Biografia) > 0 {
		register["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		register["ubicacion"] = u.Ubicacion
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.SitioWeb) > 0 {
		register["sitioWeb"] = u.SitioWeb
	}
	register["fechaNacimiento"] = u.FechaNacimiento

	//We perform the setting of the update record
	updtString := bson.M{
		"$set": register,
	}
	//Indicate the ID of the user
	objID, _ := primitive.ObjectIDFromHex(ID)

	//Add filter for upgrade
	filter := bson.M{"_id": bson.M{"$eq": objID}} //eq=equal

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}
