package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/jwt"
	"github.com/axi93/twittgo/models"
)

//Login do Login
func Login(w http.ResponseWriter, r *http.Request) {

	//We take the header where is all content
	w.Header().Add("content-type", "application/json")

	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User and/or Password incorrect"+err.Error(), 400)
		return
	}
	//We create the validation
	if len(t.Email) == 0 {
		http.Error(w, "The email of user is required", 400)
		return
	}
	document, exits := db.TryLogin(t.Email, t.Password)
	if exits == false {
		http.Error(w, "User and/or Password incorrect", 400)
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error ocurred with the Token"+err.Error(), 400)
	}

	resp := models.AnswerLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//How to save a Cookie this json
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
