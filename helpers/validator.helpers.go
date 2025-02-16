package helpers

import "github.com/go-playground/validator/v10"

func Validator(data interface{}) error {

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(data)

	return err
}
