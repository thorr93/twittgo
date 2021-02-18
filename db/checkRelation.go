package db

import (
	"context"
	"fmt"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckRelation check the relation between 2 users
func CheckRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("twittgo")
	col := bd.Collection("relation")

	//a search condition is created
	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	//If it have a result, we make a Decode to result
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil

}
