package auth

import (
	"app/models"
)

type AuthSuccessResponse struct {
	Token string `json:"token"`
	User  *models.User `json:"user"`
}
