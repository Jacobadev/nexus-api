package handler

import (
	"net/http"

	"github.com/gateway-address/user"
)

func ExtractUserInput(r *http.Request) user.User {
	user := user.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		UserName:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}
	return user
}
