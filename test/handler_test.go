package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gateway-address/handler"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetUsersHandler(t *testing.T) {
	// Create a request to simulate GET /user
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly with the mock user handler
	handler := handler.GetUsersHandler(userHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type header
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	// Check the response body
	expectedResponseBody := `[{"ID":1,"FirstName":"John","LastName":"Doe","UserName":"johndoe","Email":"johndoe@example.com"},{"ID":2,"FirstName":"Jane","LastName":"Smith","UserName":"janesmith","Email":"janesmith@example.com"}]`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponseBody)
	}
}
