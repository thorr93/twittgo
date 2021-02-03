package middlew

import (
	"net/http"

	"github.com/axi93/twittgo/db"
)

//SE TRAEN LAS CONEXIONES QUE ESTAN EN DB

/*CheckDB - Es el middlew que permite conocer el estado de la BD. Comprueba la DB y trae las conexiones que estan en DB.
 Recive una funcion i envia una funcion
Recibe la conexion y la tengo que enviar tal cual */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Connection lost with DataBase", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
