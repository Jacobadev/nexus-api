package auth

import "net/http"

type Handlers interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	GetUserByID() http.HandlerFunc
	GetUsers() http.HandlerFunc
}
