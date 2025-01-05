package writer

import (
	"errors"

	"github.com/Alitindrawan24/go-currency-conversion/internal/pkg/exception"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func APIResponse(message string, status string, data any) Response {
	jsonResponse := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	if status == "fail" {
		jsonResponse.Data = nil
		jsonResponse.Error = data
	}

	return jsonResponse
}

func APIErrorResponse(message string, err error) Response {

	var exception *exception.Exception
	if errors.As(err, &exception) {
		message = err.Error()
	}

	jsonResponse := Response{
		Status:  "fail",
		Message: message,
	}

	return jsonResponse
}
