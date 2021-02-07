package routers

import (
	"encoding/json"
	"net/http"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//ModifyProfile - modify the profile of the user
func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Inccorrect data"+err.Error(), 400)
		return
	}

	var status bool
	//err ya lo tengo inicialicado por eso le pongo el igual
	status, err = db.ModifyRegister(t, IDUser)
	if err != nil {
		http.Error(w, "An error ocurred when we try to modify the register. Try again"+err.Error(), 400)
		return
	}
	//No hay error pero tampoco se ha modificado
	if status == false {
		http.Error(w, "Imposible modify the register of the user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
