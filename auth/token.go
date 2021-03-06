package auth

import (
	"log"
	"time"
	"app/models"

	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "some-secret-key"
	SigningMethod = "HS256"
)

func CreateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(SigningMethod))
	claims := make(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["userEmail"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("error parsing token")
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(validToken string) (interface{}, error){
	token, err := jwt.Parse(validToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Printf("invalid token")
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"], nil
	}
	return nil, nil
}
