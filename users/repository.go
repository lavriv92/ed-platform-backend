package users

import (
	"log"
	"app/models"
)

func FindByEmail(email string) (*models.User, error){
	var user models.User
	session, err := models.NewSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = session.Collection("users").Find().Where("email = ?", email).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindById(id int64) (*models.User, error) {
	var user models.User
	session, err := models.NewSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = session.Collection("users").Find().Where("id = ?", id).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}