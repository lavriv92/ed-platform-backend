import (
	"app/models"
)

func FindByEmail() {
	session, err := models.Session()
	session.Collection("users").Find().All()
}