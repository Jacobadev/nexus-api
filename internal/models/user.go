package model

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id" db:"user_id" validate:"omitempty"`
	UserName  string    `json:"username" db:"username" validate:"required,gte=5,lte=20"`
	Password  string    `json:"password,omitempty" db:"password" validate:"required,gte=8,containsany=!@#$,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`
	Email     string    `json:"email,omitempty" db:"email" validate:"required,omitempty,lte=60,email"`
	FirstName string    `json:"first_name" db:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName  string    `json:"last_name" db:"last_name" redis:"last_name" validate:"required,lte=30"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

// Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Prepare user for register
func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	return nil
}

// Prepare user for register
func (u *User) PrepareUpdate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	return nil
}

type UsersList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
