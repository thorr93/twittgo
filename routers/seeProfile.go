package routers

import (
	"encoding/json"
	"net/http"

	"github.com/axi93/twittgo/db"
)

//SeeUser - Let extract the data from profile
func SeeUser(w http.ResponseWriter, r *http.Request) {
	//Check if we have and ID
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}
	//Profile not find
	profile, err := db.SearchUser(ID)
	if err != nil {
		http.Error(w, "An error occurred when we try to find the register"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
