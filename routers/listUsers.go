package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/axi93/twittgo/db"
)

//ListUsers read the list of users
func ListUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("stype")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Must send the parameter page as a enter major to 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
