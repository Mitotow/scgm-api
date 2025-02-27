package models

import (
	"github.com/Mitotow/scgm-api/config"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

var messages *config.Messages = config.GetMessages()

func CreateInternalServerError() *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  messages.InternalServerError,
	}
}
