package user

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll() ([]interface{}, error)
	Create(userInfo User) error
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

func NewUser() *User {
	user := User{
		ID:        int(uuid.New().ID()),
		CreatedAt: "2022-11-11T11:11:11Z",
		UpdatedAt: "2022-11-11T11:11:11Z",
	}
	return &user
}
