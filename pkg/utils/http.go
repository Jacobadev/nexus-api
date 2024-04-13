package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gateway-address/config"
	"github.com/gorilla/mux"
)

func ReadRequest(r *http.Request, u interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, u); err != nil {
		return fmt.Errorf("error Unmarshalling: %v", err)
	}

	return nil
}

// Get request id from echo context
func GetRequestID(r *http.Request) (int, error) {
	// Obtaining the URL parameters from the request
	paramVars := mux.Vars(r)

	// Checking if the 'id' parameter exists in the URL parameters
	idString, ok := paramVars["id"]
	if !ok {
		return 0, fmt.Errorf("ID not found in request parameters")
	}

	// Parsing the ID string to an integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID: %v", err)
	}

	// Returning the ID and any error
	return id, nil
}

// ReqIDCtxKey is a key used for the Request ID in context

func ConfigureJWTCookie(cfg *config.Config, jwtToken string) *http.Cookie {
	return &http.Cookie{
		Name:       cfg.Cookie.Name,
		Value:      jwtToken,
		Path:       "/",
		RawExpires: "",
		MaxAge:     cfg.Cookie.MaxAge,
		Secure:     cfg.Cookie.Secure,
		HttpOnly:   cfg.Cookie.HTTPOnly,
		SameSite:   0,
	}
}

func GetIPAddress(r *http.Request) string {
	return r.RemoteAddr
}

func CreateSessionCookie(cfg *config.Config, session string) *http.Cookie {
	return &http.Cookie{
		Name:  cfg.Session.Name,
		Value: session,
		Path:  "/",
		// Domain: "/",
		// Expires:    time.Now().Add(1 * time.Minute),
		RawExpires: "",
		MaxAge:     cfg.Session.Expire,
		Secure:     cfg.Cookie.Secure,
		HttpOnly:   cfg.Cookie.HTTPOnly,
		SameSite:   0,
	}
}

func DeleteSessionCookie(w http.ResponseWriter, sessionName string) {
	cookie := &http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

// Error response with logging error for echo context

// Read sanitize and validate request
