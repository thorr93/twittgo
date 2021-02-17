package routers

import (
	"net/http"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//DeleteRelation realise the delete of the relation between users
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The parameter ID is necesary", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "An error occured trying to delete the relation"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Relation can't be deleted"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
