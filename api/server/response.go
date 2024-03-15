package server

import (
	"encoding/json"
	"net/http"
)

func WriteUserPOSTResponse(w http.ResponseWriter) {
	var data interface{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func WriteUserGETResponse(w http.ResponseWriter, usersFromDB interface{}) {
	users, err := json.Marshal(usersFromDB)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(users)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
