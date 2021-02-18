package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/axi93/twittgo/db"
)

//ReadTweetsFollowers read the tweets of all our followers
func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send the parameter page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Must send the parameter page as enter major to 0", http.StatusBadRequest)
		return
	}
	answer, correct := db.ReadTweetsFollowers(IDUser, page)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
