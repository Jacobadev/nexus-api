package utils

import (
	"errors"
	"fmt"

	model "github.com/gateway-address/internal/models"
	"github.com/go-playground/validator/v10"
)

func ValidateUser(user *model.User) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		var errMsg string
		for _, e := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("Tag: '%s', Namespace: '%s', Type: %s,  Value: %v, Align: %d, Param: %s \n",
				e.Tag(), e.StructNamespace(), e.Type(), e.Value(), e.Type().Align(), e.Param())
			fmt.Println(errMsg)
			fmt.Println(err.(validator.ValidationErrors))
		}
		return errors.New(errMsg)
	}
	return nil
}
