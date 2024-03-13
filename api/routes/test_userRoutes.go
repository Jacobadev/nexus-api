package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gateway-address/api/config"
)

func test_get_all_users(t *testing.T) {
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	baseURL := fmt.Sprintf("http://localhost:%d", cfg.Server.Port)
	req := httptest.NewRequest("GET", baseURL, nil)
	w := httptest.NewRecorder()
	UserGetAll(w, req)
	response := w.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK || err != nil {
		t.Errorf(`Expected: %d,Received: %d, Error: %s`, http.StatusOK, response.StatusCode, err)
	}
}
