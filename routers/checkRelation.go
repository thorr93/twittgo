package routers

import (
	"encoding/json"
	"net/http"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//CheckRelation realise the delete of the relation between users
func CheckRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var answer models.AnswerCheckRelation

	status, err := db.CheckRelation(t)
	if err != nil || status == false {
		answer.Status = false
	} else {
		answer.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)

}
