package validate

import (
	"errors"

	model "github.com/gateway-address/internal/models"
)

// Validator is the interface for validating user input.
type Validator interface {
	ValidateUserName(username string) error
	ValidateEmail(email string) error
	ValidatePassword(password string) error
	ValidateFirstName(firstName string) error
	ValidateLastName(lastName string) error
}

// UserValidator is a concrete implementation of the Validator interface.
type UserValidator struct{}

// NewValidator creates a new instance of UserValidator.
func NewValidator() *UserValidator {
	return &UserValidator{}
}

// CreateUserInput validates user input for creating a new user.
func (uv *UserValidator) ValidateUser(user model.User) error {
	// Validate user input
	if err := uv.ValidateUsernameLength(user.UserName); err != nil {
		return err
	}
	// Return an error if validation fails
	return nil
}

const (
	MINIMUM_USERNAME_LENGTH = 6
	MAXIMUM_USERNAME_LENGTH = 32
)

const (
	EMPTY = 0
)

const (
	MINIMUM_PASSWORD_LENGTH = 6
	MAXIMUM_PASSWORD_LENGTH = 32
)

func (uv *UserValidator) ValidateUsernameLength(username string) error {
	usernameLength := len(username)
	if usernameLength == EMPTY {
		return errors.New("username cannot be empty")
	}

	if usernameLength < MINIMUM_USERNAME_LENGTH {
		return errors.New("username too short")
	}
	if usernameLength > MAXIMUM_USERNAME_LENGTH {
		return errors.New("Username too long")
	}
	return nil
}

func ValidateEmail(email string) error {
	return nil
}

func ValidatePassword(password string) error {
	return nil
}

func ValidateFirstName(firstName string) error {
	return nil
}

func ValidateLastName(lastName string) error {
	return nil
}
