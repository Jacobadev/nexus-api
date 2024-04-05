package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type UserRepository interface {
	GetAll() ([]interface{}, error)
	GetById(id int) (interface{}, error)
	Create(userInfo User) (sql.Result, error)
	UpdateByID(id int, userInfo User) (sql.Result, error)
	DeleteByID(id int) (sql.Result, error)
	PartialUpdateByID(id int, userInfo User) (sql.Result, error)
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ExtractUserInput(r *http.Request) (*User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Create a new user instance with default values
	newUser := &User{}

	// Unmarshal request body JSON data into newUser
	err = json.Unmarshal(body, newUser)
	if err != nil {
		return nil, err
	}
	if newUser.FirstName == "" {
		return nil, errors.New("first_name is required")
	}
	if newUser.LastName == "" {
		return nil, errors.New("last_name is required")
	}
	if newUser.UserName == "" {
		return nil, errors.New("username is required")
	}
	if newUser.Email == "" {
		return nil, errors.New("email is required")
	}
	if newUser.Password == "" {
		return nil, errors.New("password is required")
	}

	return newUser, nil
}
