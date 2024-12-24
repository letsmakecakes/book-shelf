package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

const alphaSpaceRegexString string = "^[a-zA-Z\\s]+$"

// isAlphaSpace is a custom validation function that checks if a string contains only alphabetic characters and spaces.
func isAlphaSpace(fl validator.FieldLevel) bool {
	alphaSpaceRegex := regexp.MustCompile(alphaSpaceRegexString)
	return alphaSpaceRegex.MatchString(fl.Field().String())
}

// New creates and returns a new validator instance with custom validations and tag name functions.
func New() *validator.Validate {
	validate := validator.New()

	// RegisterTagNameFunc customizes the tag name extraction for JSON tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// RegisterValidation adds the custom alphaspace validation rule.
	err := validate.RegisterValidation("alphaspace", isAlphaSpace)
	if err != nil {
		return nil
	}

	return validate
}
