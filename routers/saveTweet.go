package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//SaveTweet allow us ti save the tweet into DB
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "An error occurred when we try to insert the register, try again"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "An error occurred when we try to insert the register, try again", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
