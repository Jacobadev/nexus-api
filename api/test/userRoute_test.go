package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gateway-address/handler"
	"github.com/gateway-address/model"
)

func TestGetAllUsers(t *testing.T) {
	// Mock HTTP request
	req := httptest.NewRequest("GET", "http://localhost:3333/user", nil)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Call the handler function
	user := handler.UserGetAll()

	// Get the HTTP response
	response := w.Result()

	// Check if the status code is OK
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, received %d", http.StatusOK, response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	fmt.Println("Response Body:", string(body))
	// Validate the response body contains valid JSON
	var users []model.User
	if err := json.Unmarshal(body, &users); err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
		return
	}

	// Optionally, validate the contents of the users slice
	if len(users) == 0 {
		t.Error("Expected non-empty list of users, but received empty list")
		return
	}
}
