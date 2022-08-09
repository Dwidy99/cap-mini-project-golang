package handler

import (
	"errors"
	"fmt"
	"funding-api/campaign"
	"funding-api/helper"
	"funding-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func getCurrentUserJWT(c *gin.Context) int {
	currentUser := c.MustGet("current_user").(user.User)
	return currentUser.ID
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID := getCurrentUserJWT(c)

	pagin := helper.GeneratePaginationRequest(c)
	// users, errUser := user.Repository.FindById(userID)

	campaigns, err := h.service.GetCampaigns(userID, *pagin)
	if err != nil {
		response := helper.APIResponse("Error to Get Data Campaigns", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// user := user.User{}
	// user.Email =
	fmt.Println("CAMPAIGNS", campaigns) 
	response := helper.APIResponse("Data of Campaigns", http.StatusOK, "Success", campaigns)
	c.JSON(http.StatusOK, response)
}

// tnagkap paramaeter di handler
// handler ke service
// service yg menentukan repository mana yg di-call
// repository: FindAll, FindByUserID
// db

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// api/v1/campaigns/2
	// handler : mapping id yg ada di url ke struct input => service, call formatter
	// service : input struct input => menagkap id di url, memanggil repo
	// repository get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to Get Detail of Campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		messError := fmt.Sprintf("campaign with id %v is empty", input.ID)
		response := helper.APIResponse(messError, http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("Campaign Detail", http.StatusOK, "Success", campaignDetail)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) DeleteCampaign(c *gin.Context) {
	// campaignID := c.Params("id")
	// campaignId, err := strconv.Atoi(campaignID)
	var inputID campaign.GetCampaignDetailInput
	// var userId campaign.CreateCampaignInput
	
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete Campaign", http.StatusBadRequest, "Error", nil)
		
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	userID := getCurrentUserJWT(c)
	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("Forbidden")}
		response := helper.APIResponse("Failed Delete Campaign", http.StatusForbidden, "Error", errorMessage)

		c.JSON(http.StatusForbidden, response)
		return
	}
	
	_, err = h.service.DeleteCampaign(inputID.ID, userID)
	// campaignDetail, err := h.service.GetCampaignByID(inputID)
	
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed Delete Campaign", http.StatusForbidden, "Error", errorMessage)

		c.JSON(http.StatusForbidden, response)
		return
	}

	response := helper.APIResponse("Success Delete Campaign", http.StatusOK, "Success", nil)

	c.JSON(http.StatusOK, response)
}

// tangkap parameter dari user input struct
// ambil current user dari jwt/handler
// panggil service, parameternya input struct (dan juga buat slug)
// panggil repository untuk simpan data campaign baru

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed Create Campaign", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed Create Campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to Create Campaign", http.StatusOK, "Success", newCampaign)
	c.JSON(http.StatusOK, response)
}

// User masukkan input
// handler
// mapping dari input ke input struct (ada 2)
// input dari user, dan juga input yg ada di url (passing ke server)
// service 
// repository update data campaign

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput
	
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed Update Campaign", http.StatusBadRequest, "Error", nil)
		
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	var inputData campaign.CreateCampaignInput
	
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed Update Campaign", http.StatusBadRequest, "Error", errorMessage)
		
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	currentUser := c.MustGet("current_user").(user.User)
	inputData.User = currentUser
	
	updateCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	fmt.Println("Tracking Error: ", err)
	
	if err != nil {
		errorMessage := gin.H{"error": nil}
		response := helper.APIResponse("Failed Update Campaign", http.StatusBadRequest, "Error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Update Campaign", http.StatusOK, "Success", updateCampaign)

	c.JSON(http.StatusBadRequest, response)
}

// handler
// tangkap input dan ubah ke struct input
// save image campaign ke suatu folder 
// service (kondisi memanggil point 2 di repository point 2)
// repository :
// 1. create image/save data image ke dalam tabel campaign_image
// 2. ubah is_primary true ke false

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput
	
	err := c.ShouldBind(&input)
	if err != nil {
		data := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": data}

		response := helper.APIResponse("Failed Update Campaign", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload Campaign Image", http.StatusBadRequest, "Error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload Campaign Image", http.StatusBadRequest, "Error", data)
	
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload Campaign Image", http.StatusBadRequest, "Error", data)
	
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Campaign Image Successfully Uploded", http.StatusOK, "Success", data)

	c.JSON(http.StatusOK, response)
}

