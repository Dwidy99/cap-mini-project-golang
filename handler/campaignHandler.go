package handler

import (
	"funding-api/campaign"
	"funding-api/helper"
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