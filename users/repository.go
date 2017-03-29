package users

import (
	"log"
	"app/models"
)

const (
	CollectionName = "users"
)

func FindByEmail(email string) (*models.User, error){
	var user models.User
	session, err := models.NewSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = session.Collection(CollectionName).Find().Where("email = ?", email).One(&user)
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
	err = session.Collection(CollectionName).Find().Where("id = ?", id).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(user models.User) error {
	session, err := models.NewSession()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = session.Collection(CollectionName).Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}