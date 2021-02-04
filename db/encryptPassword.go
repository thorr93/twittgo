package db

import (
	"golang.org/x/crypto/bcrypt"
)

//EncryptPassword it's done for encrypt de password. Price is an algrothim and it's de number of times is encrypted the password. In this case 2*8
func EncryptPassword(pass string) (string, error) {
	price := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), price)

	return string(bytes), err

	//return bcrypt.GenerateFromPassword([]byte(pass), price)
}
