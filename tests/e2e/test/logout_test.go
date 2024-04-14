package e2e

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	model "github.com/gateway-address/internal/models"
	testsetup "github.com/gateway-address/tests/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LogoutTestSuite struct {
	suite.Suite
	URL  string
	user *model.User
}

func (s *LogoutTestSuite) SetupSuite() {
	s.URL = testsetup.AuthURL + "logout"
	s.user = testsetup.UserGenerator()
}

func (s *LogoutTestSuite) TestLogoutEndpointWorking() {
	res, _, err := testsetup.LoginSetupResponse(s.T())
	if err != nil {
		s.T().Errorf("Erro ao enviar a solicitação POST: %v", err)
		return
	}
	defer res.Body.Close()
	user := &model.UserWithToken{}
	body, err := io.ReadAll(res.Body)
	json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		s.T().Errorf("Erro ao decodificar a resposta do login: %v", err)
		return
	}

	logoutReq, err := http.NewRequest("POST", s.URL, nil)
	if err != nil {
		s.T().Errorf("Erro ao criar a requisição de logout: %v", err)
		return
	}
	logoutReq.AddCookie(&http.Cookie{Name: "session-id", Value: user.Token})

	client := http.Client{}
	logoutRes, err := client.Do(logoutReq)
	if err != nil {
		s.T().Errorf("Erro ao enviar a solicitação de logout: %v", err)
		return
	}
	defer logoutRes.Body.Close()

	// Verifique se o logout foi bem-sucedido (esperando StatusCode 200)
	assert.Equal(s.T(), http.StatusOK, logoutRes.StatusCode)
}

func (s *LogoutTestSuite) TestLogoutUnauthorized() {
	c := http.Client{}
	res, err := c.Post(s.URL, "application/json", nil)
	if err != nil {
		s.T().Errorf("Erro ao enviar a solicitação POST: %v", err)
		return
	}
	defer res.Body.Close()

	assert.Equal(s.T(), http.StatusUnauthorized, res.StatusCode)
}

func TestLogoutEndpointSuite(t *testing.T) {
	suite.Run(t, new(LogoutTestSuite))
}
