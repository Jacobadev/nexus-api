package user

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll() ([]interface{}, error)
	Create(userInfo User) (sql.Result, error)
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ExtractUserInput(r *http.Request) (*User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Create a new user instance with default values
	newUser := &User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Unmarshal request body JSON data into newUser
	err = json.Unmarshal(body, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
