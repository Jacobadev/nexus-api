package routes

import (
	"net/http"

	"github.com/gateway-address/api/controller"
	"github.com/gorilla/mux"
)

func ValidateUserInput(r *http.Request) error {
	return nil
}

func RegisterUserHandler(mux *mux.Router) {
	mux.HandleFunc("/user", controller.UserGETController).Methods("GET")
	mux.HandleFunc("/user", controller.UserPOSTController).Methods("POST")
}
