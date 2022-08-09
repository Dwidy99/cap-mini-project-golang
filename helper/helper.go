package helper

import (
	"funding-api/campaign"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func GeneratePaginationRequest(context *gin.Context) *campaign.Pagination {
	// convert query parameter string to int
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "0"))
	return &campaign.Pagination{Limit: limit, Page: page}
}