package server

import (
	"encoding/json"
	"net/http"
)

func WriteUserPOSTResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func WriteUserGETResponse(w http.ResponseWriter, users interface{}) {
	usersJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(usersJson)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
