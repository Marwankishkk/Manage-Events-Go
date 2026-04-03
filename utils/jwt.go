package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = "mysecretkey"

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 72).Unix()})

	return token.SignedString([]byte(secret))

}
