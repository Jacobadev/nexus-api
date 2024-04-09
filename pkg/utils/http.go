package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gateway-address/config"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/logger"
)

func ReadRequest(r *http.Request) (*model.User, error) {
	fmt.Println("Reading Request...")
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get request id from echo context
func GetRequestID(r *http.Request) string {
	return r.Context().Value("id").(string)
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

// Error response with logging error for echo context
func LogResponseError(r *http.Request, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(r),
		GetIPAddress(r),
		err,
	)
}

// Read sanitize and validate request
