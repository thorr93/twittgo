package db

import (
	"context"
	"fmt"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadAllUsers - Read all users recorded into the sistem, if we recive "R" in which it gives us, just only record what is relationated with us
func ReadAllUsers(ID string, page int64, search string, stype string) ([]*models.Users, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	//Aqui apunto a la base de datos y la collection
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("users")

	var results []*models.Users

	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip((page - 1) * 20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var find, include bool

	for cur.Next(ctx) {
		var s models.Users
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		find, err = CheckRelation(r)
		if stype == "new" && find == false {
			include = true
		}
		if stype == "follow" && find == true {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}
		if include == true {
			s.Password = ""
			s.Biografia = ""
			s.Banner = ""
			s.Email = ""
			s.SitioWeb = ""
			s.Ubicacion = ""

			results = append(results, &s)
		}

	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
