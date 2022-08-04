package main

import (
	"fmt"
	"funding-api/handler"
	"funding-api/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=d dbname=crowd_founding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	
	userRepository := user.NewRepositoryDB(db)
	userService := user.NewService(userRepository)

	userByEmail, err := userRepository.FindByIdEmail("golang@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	
	if userByEmail.ID == 0 {
		fmt.Println("User tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.RegisterUser)

	router.Run()

}