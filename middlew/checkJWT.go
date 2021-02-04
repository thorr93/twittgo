package middlew

import (
	"net/http"

	"github.com/axi93/twittgo/routers"
)

//ValidJWT allows validate the JWT that come to us as a request
func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error into the Token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
