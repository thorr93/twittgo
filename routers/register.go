package routers

import (
	"encoding/json"
	"net/http"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//Register is the function to create the user record in the DB.
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.Users
	//first I create a model of the user and then decode it, then everything that comes in body will be a json-like structure that we have already created.
	err := json.NewDecoder(r.Body).Decode(&t) //The Body object is a stream, i.e. it is only used once and then destroyed.
	if err != nil {
		http.Error(w, "Error into the data recibed"+err.Error(), 400)
		return
	}
	//Validations- email equal 0
	if len(t.Email) == 0 {
		http.Error(w, "The email of user is required", 400)
		return
	}
	//Validations - password <6
	if len(t.Password) < 6 {
		http.Error(w, "The Password must be 6 characters or more", 400)
		return
	}
	//Validations - check if email exist
	_, find, _ := db.CheckStillUser(t.Email)
	if find == true {
		http.Error(w, "There's exist a user with that email", 400)
		return
	}
	//Validations - it is checked for proper registration
	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error ocurred at the time of register the error, please try it again"+err.Error(), 400)
		return
	}
	//Validations - check again that it has been registered correctly (problems with mongodb that sometimes does not return true)
	if status == false {
		http.Error(w, "An error ocurred at the time of register the error, maybe its mongoDB", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
