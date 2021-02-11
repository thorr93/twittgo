package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/axi93/twittgo/db"
)

//ReadTweets - Read the Tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must sent the parameter ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must sent the parameter page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Must sent the parameter page with value more than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	answer, correct := db.ReadTweet(ID, pag)
	if correct == false {
		http.Error(w, "Error when we reed tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
