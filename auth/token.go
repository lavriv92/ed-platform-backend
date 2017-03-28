package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "some-secret-key"
	SigningMethod = "HS256"
)

func CreateToken(id uint64) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(SigningMethod))
	claims := make(jwt.MapClaims)
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("error parsing token")
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(validToken string) (string, error){
	token, err := jwt.Parse(validToken, func (token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil && !token.Valid {
		log.Printf("invalid token")
		return "", err
	}
	return "", nil
}
