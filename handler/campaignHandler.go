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

	campaign, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to Get Data Campaigns", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data of Campaigns", http.StatusOK, "Success", campaign)
	c.JSON(http.StatusOK, response)
}

// tnagkap paramaeter di handler
// handler ke service
// service yg menentukan repository mana yg di-call
// repository: FindAll, FindByUserID
// db