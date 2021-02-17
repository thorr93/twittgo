package db

import (
	"context"
	"log"
	"time"

	"github.com/axi93/twittgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadTweet allows reads a tweet from a user
func ReadTweet(ID string, page int64) ([]*models.ReturnTweets, bool) { //devuelvo todos los tweets de un usuario a la vez en un slice (array)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	bd := MongoCN.Database("twittgo")
	col := bd.Collection("tweet")

	//Donde grabamos los resultados
	var result []*models.ReturnTweets

	condition := bson.M{
		"userID": ID,
	}

	//Le doy la primera pagina, le resto -1 pero al ser pagina 1 es 1-1=0*20=0, por tanto coge los 20 primeros tweets. Segunda hoja 2-1*20=20, Skipea los primeros 20 tweets
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((page - 1) * 20)

	pointer, err := col.Find(ctx, condition, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.ReturnTweets
		err := pointer.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}
