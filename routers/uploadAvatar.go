package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
)

//UploadAvatar - It's used to Upload the Banner
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var doc string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(doc, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error when we upload the image"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error when we copy the image"+err.Error(), http.StatusBadRequest)
		return
	}

	var users models.Users
	var status bool
	users.Avatar = IDUser + "." + extension
	status, err = db.ModifyRegister(users, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error when we save the avatar into DB"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
