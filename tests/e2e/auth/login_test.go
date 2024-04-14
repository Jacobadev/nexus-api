package authtest

import (
	"encoding/json"
	"net/http"
	"testing"

	model "github.com/gateway-address/internal/models"
	testsetup "github.com/gateway-address/tests/e2e"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	URL  string
	user *model.User
}

func (s *LoginTestSuite) SetupSuite() {
	s.URL = testsetup.AuthURL + "login"
	s.user = testsetup.UserGenerator()
}

func (s *LoginTestSuite) TestLoginShouldReturnOK() {
	res, user, err := testsetup.LoginSetupResponse(s.T())
	if err != nil {
		s.T().Errorf("Erro ao enviar a solicitação POST: %v", err)
		return
	}
	defer res.Body.Close()

	// Verifique se a resposta tem status OK (200)
	assert.Equal(s.T(), http.StatusOK, res.StatusCode)
	// Decodifique o JSON da resposta
	var responseBody struct {
		User  model.User `json:"user"`
		Token string     `json:"token"`
	}

	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		s.T().Errorf("Erro ao decodificar o JSON da resposta: %v", err)
		return
	}
	s.NotEmpty(s.T(), responseBody.Token)
	assert.Equal(s.T(), user.Email, responseBody.User.Email)
}

func TestLoginEndpointSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
