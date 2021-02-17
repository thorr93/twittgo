package db

import (
	"context"
	"time"

	"github.com/axi93/twittgo/models"
)

//DeleteRelation - Delete the relation into the DB
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	//Aqui apunto a la base de datos y la collection
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
