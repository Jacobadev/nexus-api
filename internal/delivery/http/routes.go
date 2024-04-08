package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MapAuthRoutes(h *authHandlers, r *mux.Router) {
	r.HandleFunc("/signup", h.Register()).Methods("POST")
	// r.HandleFunc("/", h.GetUsers()).Methods("GET")
	// r.HandleFunc("/{id}", h.GetUserByID()).Methods("GET")
	// r.HandleFunc("/{id}", h.Delete()).Methods("DELETE")
	// r.HandleFunc("/{id}", h.UpdateByID()).Methods("PUT")
	// r.HandleFunc("/{id}", h.PartialUpdateUserByID()).Methods("PATCH")
}

func HealthCheck(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
