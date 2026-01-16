package handlvalidator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	emailTag    = "email"
	requiredTag = "required"
)

var (
	emailTagErr    = "invalid email"
	requiredTagErr = "field is required"
	defaultErr     = "field is invalid"
)

func MapValidationErrors(err error) (map[string]string, bool) {
	var vl validator.ValidationErrors
	if errors.As(err, &vl) {
		errMap := make(map[string]string)

		for _, fl := range vl {
			field := fl.Field()
			tag := fl.Tag()

			errMap[field] = validateError(tag)
		}

		return errMap, true
	}
	return nil, false
}

func validateError(tag string) string {
	var res string
	switch tag {
	case emailTag:
		res = emailTagErr
	case requiredTag:
		res = requiredTagErr
	default:
		res = defaultErr
	}
	return res
}
