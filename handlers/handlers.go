package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/axi93/twittgo/middlew"
	"github.com/axi93/twittgo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handling - When you call the API enter here and define the routes that we will handle
func Handling() {
	/*CIt captures the http and it gives a handling to the respons writer and to the request that comes
	in the call of the API sive for if in the body of the call there is information and it will send the answer to the navigator.
	When we call the route of the API, this same one returns a status.
	That is to say it gives and receives information*/
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/seeProfile", middlew.CheckDB(middlew.ValidJWT(routers.SeeProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlew.CheckDB(middlew.ValidJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweet", middlew.CheckDB(middlew.ValidJWT(routers.ReadTweet))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadavatar", middlew.CheckDB(routers.UploadAvatar)).Methods("POST")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(routers.UploadBanner)).Methods("POST")
	router.HandleFunc("/obtainAvatar", middlew.CheckDB(routers.ObtainAvatar)).Methods("POST")
	router.HandleFunc("/obtainBanner", middlew.CheckDB(routers.ObtainBanner)).Methods("POST")

	router.HandleFunc("/registerRelation", middlew.CheckDB(middlew.ValidJWT(routers.RegisterRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlew.CheckDB(middlew.ValidJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/checkRelation", middlew.CheckDB(middlew.ValidJWT(routers.CheckRelation))).Methods("GET")

	router.HandleFunc("/listUsers", middlew.CheckDB(middlew.ValidJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/readTweetsFollowers", middlew.CheckDB(middlew.ValidJWT(routers.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	//We give permissions to the API to anyone so that it can work without any problem.
	handler := cors.AllowAll().Handler(router)
	//Sets the server to listen to the port
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
