package user

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (i *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupasion = input.Occupasion
	password_hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal("Password Hash Crash!! \n")
	}
	user.Password_hash = string(password_hash)
	user.Role = "user"
	user.Token = "sasasasasa"
	
	newUser, err := i.repository.Save(user)
	if err != nil {
		log.Fatal("New User Not Created!!")
		return newUser, err
	}
	fmt.Println("NEW USER!!!: ", newUser)

	return newUser, nil
}