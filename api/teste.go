package 

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define the handler function for the "/ability" endpoint
	router.HandleFunc("/ability", AbilityHandler).Methods("GET")

	// Start the HTTP server on port 8080
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

// AbilityHandler is the handler function for the "/ability" endpoint
func AbilityHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query parameters "limit" and "offset" from the request URL
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	// Check if both "limit" and "offset" parameters are provided
	if limit != "" && offset != "" {
		// Convert the parameters to integers
		// (Error handling is omitted for brevity)
		fmt.Printf("Limit: %s, Offset: %s\n", limit, offset)
	} else {
		// Handle the case where one or both parameters are missing
		fmt.Println("Both 'limit' and 'offset' parameters are required")
	}

	// Write a response to the client
	fmt.Fprintf(w, "AbilityHandler called\n")
}
