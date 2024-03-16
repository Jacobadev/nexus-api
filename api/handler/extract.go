package handler

import (
	"net/http"

	"github.com/gateway-address/user"
	"github.com/gorilla/mux"
)

func ExtractUserInput(r *http.Request) user.User {
	vars := mux.Vars(r)
	user := user.User{
		FirstName: vars["first_name"],
		LastName:  vars["last_name"],
		UserName:  vars["username"],
		Email:     vars["email"],
		Password:  r.FormValue("password"), // Extract password from form values
	}
	return user
}
