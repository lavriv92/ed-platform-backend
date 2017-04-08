package users

import (
	"log"
	"app/models"

	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	CollectionName = "users"
)


type UsersRepository {
	dbSession  sqlbuilder.Database
}

func (r *UsersRepository) FindByEmail(email string) (*models.User, error){
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

func (r *UsersRepository) FindById(id int64) (*models.User, error) {
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

func (r *UsersRepository) Create(user models.NewUser) error {
	session, err := models.NewSession()
	if err != nil {
		log.Fatal(err)
		return err
	}
	session.Collection(CollectionName).Insert(user)
	return nil
}