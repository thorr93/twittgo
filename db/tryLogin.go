package db

import (
	"github.com/axi93/twittgo/models"
	"golang.org/x/crypto/bcrypt"
)

//TryLogin check login into DB
func TryLogin(email string, password string) (models.Users, bool) {
	usu, find, _ := CheckStillUser(email)
	if find == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	//Check password from parametrer and from database. Databse Pass is encrypted so decrypt.
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true

}
