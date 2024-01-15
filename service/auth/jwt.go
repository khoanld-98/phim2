package auth

import (
	"os"
	"web/models"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(User *models.Users) (string, error) {
	mySigningKey := []byte(os.Getenv("SIGNING_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"User": User,
	})

	return token.SignedString(mySigningKey)
}

func ParserJwt(tokenJwt string) *jwt.Token {
	token, _ := jwt.Parse(tokenJwt, func(toke *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	return token
}
