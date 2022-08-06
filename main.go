package main

import (
	"fmt"
	"funding-api/auth"
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
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMH0.DAZv9a5nwrom1zKf_n6A4CjHkms77L7u01eBvxui78")
	if err!= nil {
		fmt.Println("ERRROR")
		fmt.Println("ERRROR")
	}

	if token.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {
		fmt.Println("NOT VALID")
	}

	userHandler := handler.NewUserHandler(userService, authService)


	router := gin.Default()
	api := router.Group("/api/v1/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/emailChecker", userHandler.CheckEmailIsAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()

}