package user

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type UserRepository interface {
	GetAll() ([]interface{}, error)
	GetById(id int) (interface{}, error)
	Create(userInfo User) (sql.Result, error)
	Update(id int, userInfo User) (sql.Result, error)
	Delete(id int) (sql.Result, error)
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

	return newUser, nil
}
