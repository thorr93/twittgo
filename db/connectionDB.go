package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN -Es el objeto de conexion a la DB - Exportada (Publica) porque es mayuscula
var MongoCN = ConnectDB()

//Uso interno
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:Lv3s5olo@twitter.9ehba.mongodb.net/<dbname>?retryWrites=true&w=majority")

//ConnectDB - me permite conectar con al DB - Devuelve objeto de tipo mongo.Client, es decir devuelve una conexion
func ConnectDB() *mongo.Client {
	//Hace conexion a la base de datos con el clientOptions que ya tiene URL
	//Contextos sirven para comunicar informacion entre ejecucion i ejecucion i permite seter una serie de valores
	//Ejemplo: timeout
	client, err := mongo.Connect(context.Background(), clientOptions)
	//Si hay error indica el porque y devulve el objeto cliente ya que siempre debo devolver el objeto aunque este vacio
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	//Miramos si DB esta arriba, si no esta devuvlve error
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Creada con la DB")
	return client
}

//CheckConnection - Es el Ping a la DB
func CheckConnection() bool {
	err := MongoCN.Ping(context.Background(), nil)
	if err != nil {
		return false
	}
	return true
}
