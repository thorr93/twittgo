package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/axi93/twittgo/db"
)

//ObtainBanner send the Banner to the HTTP
func ObtainBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parametrer ID", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchUser(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copy image", http.StatusBadRequest)
	}

}
