package api

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("validateRegex", validateRegex)
	validate.RegisterValidation("validateDate", validateDate)
	validate.RegisterValidation("validateTime", validateTime)
}

func validateRegex(fl validator.FieldLevel) bool {
	pattern := fl.Param()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(fl.Field().String())
}

func validateDate(fl validator.FieldLevel) bool {
	format := fl.Param()

	_, err := time.Parse(format, fl.Field().String())
	return err == nil
}

func validateTime(fl validator.FieldLevel) bool {
	format := fl.Param()

	_, err := time.Parse(format, fl.Field().String())
	return err == nil
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func ValidateStruct(input interface{}) (*ValidationErrorResponse, error) {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Failed validation: %s", err.Tag())
		}
		return &ValidationErrorResponse{Errors: validationErrors}, fmt.Errorf("validation failed")
	}
	return nil, nil
}
