package handler

import (
	"fmt"
	"funding-api/campaign"
	"funding-api/helper"
	"funding-api/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to Get Data Campaigns", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data of Campaigns", http.StatusOK, "Success", campaign.FormatCampaigns(campaigns))
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
		response := helper.APIResponse("Failed to Get Detail of Campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("Campaign Detail", http.StatusOK, "Success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

// tangkap parameter dari user input struct
// ambil current user dari jwt/handler
// panggil service, parameternya input struct (dan juga buat slug)
// panggil repository untuk simpan data campaign baru

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	fmt.Println("ERRORNYA", err)
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

	response := helper.APIResponse("Success to Create Campaign", http.StatusOK, "Success", campaign.FormatCampaign(newCampaign))
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