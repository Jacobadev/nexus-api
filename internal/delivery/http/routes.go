package http

import "github.com/gorilla/mux"

func MapAuthRoutes(h *authHandlers, r *mux.Router) {
	r.HandleFunc("/", h.Register()).Methods("POST")
	// r.HandleFunc("/", h.GetUsers()).Methods("GET")
	// r.HandleFunc("/{id}", h.GetUserByID()).Methods("GET")
	// r.HandleFunc("/{id}", h.Delete()).Methods("DELETE")
	// r.HandleFunc("/{id}", h.UpdateByID()).Methods("PUT")
	// r.HandleFunc("/{id}", h.PartialUpdateUserByID()).Methods("PATCH")
}
