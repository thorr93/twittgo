package middlew

import (
	"net/http"

	"github.com/axi93/twittgo/db"
)

/*CheckDB - It is the middlew that allows to know the state of the DB. It checks the DB and brings the connections that are in DB.
 It receives a function and sends a function
It receives the connection and I have to send it as it is. */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Connection lost with DataBase", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
