package routers

import (
	"errors"
	"strings"

	"github.com/axi93/twittgo/db"
	"github.com/axi93/twittgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email value of Email used in all the EndPoints
var Email string

//IDUser is the ID returned from the model, it will be used in all the Endpoints
var IDUser string

//ProcessToken process token for extract their values
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("MasterofDevolopment")
	claims := &models.Claim{}
	//Hara que el token se convierta en un vector y eliminar Bearer del tokn
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Token format incorrect")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find, ID := db.CheckStillUser(claims.Email)
		if find == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, find, ID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token format Incorrect again")
	}
	return claims, false, string(""), err
}
