package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/axi93/twittgo/db"
)

//ReadTweet - Read the Tweets
func ReadTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must sent the parameter ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Must sent the parameter page", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Must sent the parameter page with value more than 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	answer, correct := db.ReadTweet(ID, pag)
	if correct == false {
		http.Error(w, "Error when we reed tweets", http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(&answer)
	if err != nil {
		log.Printf("error marshalling answer: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}
