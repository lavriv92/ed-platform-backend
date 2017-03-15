package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "some-secret-key"
)

func CreateToken() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := make(jwt.MapClaims)
	token.Claims = jwt.MapClaims{
		"userId": 1,
		"exp": tome.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("error parsing token")
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(validToken string) (string, error){
	token, err := jwt.Parse(validToken, func (token *jwt.Token) ([]byte, error) {
		return []byte(SecretKey), nil
	})
	if err != nil && !token.Valid {
		log.Printf("invalid token")
		return "", err
	}
	return token.Claims["userId"].(string)
}
