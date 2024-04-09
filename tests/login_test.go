package e2etest

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	URL string
}

func (s *LoginTestSuite) SetupTest() {
	s.URL = "http://localhost:5000/api/v1/auth/login"
}

func (s *LoginTestSuite) TestLoginShouldReturnOK() {
	const (
		// Defina os detalhes do login como uma string JSON
		loginJSON = `{
			"first_name": "Reine",
			"last_name": "Freshwater",
			"email": "rfreshwater1@godaddy.com",
			"username": "rfreshwater1",
			"password": "test_account"
		}`
	)

	// Crie um cliente HTTP
	client := http.Client{}

	// Envie a solicitação POST com os detalhes de login formatados como JSON
	res, err := client.Post(s.URL, "application/json", bytes.NewBuffer([]byte(loginJSON)))
	if err != nil {
		s.T().Errorf("Erro ao enviar a solicitação POST: %v", err)
		return
	}
	defer res.Body.Close()

	// Verifique se a resposta tem status OK (200)
	assert.Equal(s.T(), http.StatusOK, res.StatusCode)
}

func TestLoginEndpointSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
