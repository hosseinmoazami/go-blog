package errors

import (
	"blog/internal/providers/validation"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func Get() map[string]string {
	return errorsList
}
