package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterValidation("sortValidation", sortValidation)
}

func sortValidation(fl validator.FieldLevel) bool {
	sortValue := fl.Field().String()
	sortValue = strings.ToUpper(sortValue)
	return sortValue == "ASC" || sortValue == "DESC"

}
