package app

import (
	"fmt"
	"funding-api/auth"
	"funding-api/campaign"
	"funding-api/handler"
	"funding-api/helper"
	"funding-api/logger"
	"funding-api/user"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env variables")
	}
	logger.Info("Env variables run flawlessly")

	// sanityCheck()
	// db := getDBClient()
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connection")
	}
	
	userRepository := user.NewRepositoryDB(db)
	campaignRepository := campaign.NewRepositoryDB(db)
	
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/emailChecker", userHandler.CheckEmailIsAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", authMiddleware(authService, userService), campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	api.DELETE("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.DeleteCampaign)

	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer string.token.acak
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("current_user", user)

	}
}