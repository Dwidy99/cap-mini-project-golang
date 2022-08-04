package user

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByIdEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDB(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Println("Databasenya Error: ", err.Error())
		return user, nil
	}

	return user, nil
}

func (r *repository) FindByIdEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}