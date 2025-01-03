package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

// ToErrResponse converts a validation error into a structured error response.
func ToErrResponse(err error) *ErrResponse {
	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "alphaspace":
				resp.Errors[i] = fmt.Sprintf("%s can only contain alphabetic and space characters", err.Field())
			case "datetime":
				if err.Param() == "2006-01-02" {
					resp.Errors[i] = fmt.Sprintf("%s must be a valid date", err.Field())
				} else {
					resp.Errors[i] = fmt.Sprintf("%s must follow %s format", err.Field(), err.Param())
				}
			default:
				resp.Errors[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}
		return &resp
	}

	// Handle non-validation errors
	return &ErrResponse{
		Errors: []string{"an unexpected error occurred"},
	}
}
