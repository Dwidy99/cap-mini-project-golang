package handler

import (
	"fmt"
	"funding-api/auth"
	"funding-api/helper"
	"funding-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kit passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		
		response := helper.APIResponse("Failed Regester Account", http.StatusUnprocessableEntity, "ERROR", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed Regester Account", http.StatusBadRequest, "ERROR", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed Generate Token", http.StatusNetworkAuthenticationRequired, "Token Error", nil)
		c.JSON(http.StatusNetworkAuthenticationRequired, response)
		return
	}
	
	// formatter := user.FormatUser(newUser, "TokenNgasal")
	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// User memasuki input (email & password)
	// input ditangkap handler
	// mapping dari input user ke input struct

	 var input user.LoginInput
	 err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Validation Failed", http.StatusUnprocessableEntity, "ERROR", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedIn, err := h.userService.Login(input)
	
	if err != nil {
		helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Input Failed", http.StatusUnprocessableEntity, "ERROR", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedIn.ID)
	if err != nil {
		response := helper.APIResponse("Failed Generate Token", http.StatusNetworkAuthenticationRequired, "Token Error", nil)
		c.JSON(http.StatusNetworkAuthenticationRequired, response)
		return
	}
	
	// formatter := user.FormatUser(newUser, "TokenNgasal")
	formatter := user.FormatUser(loggedIn, token)

	response := helper.APIResponse("Login Successfull", http.StatusUnprocessableEntity, "SUCCESS", formatter)
	c.JSON(http.StatusOK, response)
}

func (u *userHandler) CheckEmailIsAvailability(c *gin.Context) {
	// Apakah ada inputan email dari user
	// input email di-Mapping ke struct input 
	// struct input di-passing ke service 
	// service akan memanggil repository - email sudah ada atau belum
	// repository DB

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := u.userService.IsEmailAvailable(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}

func (u *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// Simpan gambarnya di folder "images/"
	// di service kita panggil repo
	// JWT (sementara hardcode, seakan akan user yang login ID = 1)
	// repo ambil data user yg ID = 1
	// repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload File Error", http.StatusBadRequest, "Error Upload", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	currentUser := c.MustGet("current_user").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload File Failed", http.StatusBadRequest, "Error Upload", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = u.userService.SaveAvatar(userID, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload File Failed", http.StatusBadRequest, "Error upload", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar Successfuly Uploaded", http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}