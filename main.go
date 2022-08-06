package main

import (
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

	userHandler := handler.NewUserHandler(userService, authService)


	router := gin.Default()
	api := router.Group("/api/v1/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/emailChecker", userHandler.CheckEmailIsAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()

}