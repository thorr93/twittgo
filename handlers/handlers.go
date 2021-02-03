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

//Handling - Cuando llame a la API entra aqui i se definen las rutas que manejaremos
func Handling() {
	//Captura el http i le da un manejo al respons writer i al request que viene en el llamado de la API
	//sive para si en el body del llamado hay informacion y mandara la respuesta al navegador
	//Cuando llamamos la ruta de la API, esta misma devuelve un status. Es decir da i recibe informacion
	router := mux.NewRouter() //creamos un router nuevo

	//Aqui vamos creando nuestras rutas

	//cuando el en navegador se pone registro y luego va a checkDB i el return es correcto registra en routes.Register. El metodo de llamado es POST para no mostrar los datos
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	//Miramos si hay puerto y si hay lo trae, si no la seteamos
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	//Damos permisos a la API a cualquiera i que pueda funcionar sin problema
	handler := cors.AllowAll().Handler(router)
	//Pone al servidor a esuchar al puerto
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
