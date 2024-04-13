package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	model "github.com/gateway-address/internal/models"
	testsetup "github.com/gateway-address/tests/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterSuite struct {
	suite.Suite
	URL  string
	user *model.User
}

func (s *RegisterSuite) SetupSuite() {
	s.URL = testsetup.AuthURL + "register"
	s.user = testsetup.UserGenerator()
}

func (s *RegisterSuite) TestRegisterShouldReturnCreated() {
	userJson, err := json.Marshal(s.user)
	if err != nil {
		s.T().Fatal(err)
	}

	// Send POST request to the server
	c := http.Client{}
	res, err := c.Post(s.URL, "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		s.T().Fatal(err)
	}
	defer res.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		s.T().Fatal(err)
	}

	// Unmarshal response JSON into a user object
	userWIthToken := &model.UserWithToken{}
	if err := json.Unmarshal(responseBody, &userWIthToken); err != nil {
		s.T().Fatal(err)
	}

	assert.Equal(s.T(), http.StatusCreated, res.StatusCode)
	assert.NotNil(s.T(), userWIthToken.Token, "Token should not be nil")
	// Verify that the email in the response matches the email in the randomUser object
	assert.Equal(s.T(), userWIthToken.User.Email, s.user.Email)
}

//	func (s *RegisterSuite) TestShouldReturnBadRequest() {
//		c := http.Client{}
//		// Envie uma solicitação POST com um corpo JSON inválido
//		res, err := c.Post(s.URL, "application/json", bytes.NewBufferString("username: 'fodase'"))
//		if err != nil {
//			s.T().Fatal(err)
//		}
//		defer res.Body.Close()
//
//		// Verifique se o status de resposta é 400 (BadRequest)
//		assert.Equal(s.T(), http.StatusBadRequest, res.StatusCode)
//	}
func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
