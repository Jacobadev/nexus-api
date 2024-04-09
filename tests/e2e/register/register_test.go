package e2etest

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterSuite struct {
	suite.Suite
	URL string
}

func (s *RegisterSuite) SetupTest() {
	s.URL = "http://localhost:5000/api/v1/auth/register"
}

func (s *RegisterSuite) generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s *RegisterSuite) createJSON(username, email string) string {
	return fmt.Sprintf(`{
		"username": "%s",
		"password": "rxA!mp_password",
		"email": "%s",
		"first_name": "John",
		"last_name": "Doe"
	}`, username, email)
}

func (s *RegisterSuite) TestRegisterShouldReturnCreated() {
	username := "8xample_" + s.generateRandomString(8)
	email := s.generateRandomString(8) + "@businesswire.com"
	correctInfoJson := s.createJSON(username, email)

	c := http.Client{}
	res, err := c.Post(s.URL, "application/json", bytes.NewBuffer([]byte(correctInfoJson)))
	if err != nil {
		s.T().Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(s.T(), http.StatusCreated, res.StatusCode)
}

// func (s *RegisterSuite) TestShouldReturnBadRequest() {
// 	c := http.Client{}
// 	// Envie uma solicitação POST com um corpo JSON inválido
// 	res, err := c.Post(s.URL, "application/json", bytes.NewBufferString("username: 'fodase'"))
// 	if err != nil {
// 		s.T().Fatal(err)
// 	}
// 	defer res.Body.Close()
//
// 	// Verifique se o status de resposta é 400 (BadRequest)
// 	assert.Equal(s.T(), http.StatusBadRequest, res.StatusCode)
// }

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
