package jwt

import (
	"time"

	"github.com/axi93/twittgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generates the encrypt with JWT
func GenerateJWT(t models.Users) (string, error) {
	myKey := []byte("MasterofDevolopment")
	//La parte de los datos
	paylod := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	//Elegimos algormito con el que encriptara
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, paylod)
	//Le paso el String de la firma
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
