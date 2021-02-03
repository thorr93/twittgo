package routers

import (
	"encoding/json"
	"net/http"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//Register es la funcion para crear en la DB el registro de usuario
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.Users
	//Primero creo un modelo del usuario y luego lo descodifico, entonces todo lo que viene en body sera una estructira tipo json que ya tenemos creada
	err := json.NewDecoder(r.Body).Decode(&t) //El objeto Body es un stream, es decir solo se le una vez y luego se destruye.
	if err != nil {
		http.Error(w, "Error into the data recibed"+err.Error(), 400)
		return
	}
	//Se crean validaciones, si email es equal 0
	if len(t.Email) == 0 {
		http.Error(w, "The email of user is required", 400)
		return
	}
	//Se crean validaciones, si password <6
	if len(t.Password) < 6 {
		http.Error(w, "The Password must be 6 characters or more", 400)
		return
	}
	//Se crean validaciones, se revisa si el email ya existe
	_, find, _ := db.CheckStillUser(t.Email)
	if find == true {
		http.Error(w, "There's exist a user with that email", 400)
		return
	}
	//Se crean validaciones, se comprueba que se ha registrado bien
	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error ocurred at the time of register the error, please try it again"+err.Error(), 400)
		return
	}
	//Se crean validaciones,se vuelve a revisar que se ha registrado bien (problemas con mongodb que no devuelve el true a veces)
	if status == false {
		http.Error(w, "An error ocurred at the time of register the error, maybe its mongoDB", 400)
		return
	}
	//El status se crea
	w.WriteHeader(http.StatusCreated)
}
