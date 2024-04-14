package authtest

import (
	"net/http"
	"strconv"
	"testing"

	model "github.com/gateway-address/internal/models"
	testsetup "github.com/gateway-address/tests/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetUserTestSuite struct {
	suite.Suite

	URL  string
	user *model.User
}

func (s *GetUserTestSuite) SetupSuite() {
	s.URL = testsetup.AuthURL
	s.user = testsetup.UserGenerator()
}

func (s *GetUserTestSuite) TestGetUserEndpointWorking() {
	testsetup.SetupSuite(s.T(), s.user)
	user_id := strconv.Itoa(s.user.ID)
	c := http.Client{}
	res, err := c.Get(s.URL + user_id)
	if err != nil {
		s.T().Errorf("Erro ao enviar a solicitação POST: %v", err)
		return
	}
	defer res.Body.Close()
	user, err := testsetup.UserFromResponse(res)
	if err != nil {
		s.T().Errorf("Erro ao decodificar a resposta: %v", err)
		return
	}
	assert.Equal(s.T(), http.StatusOK, res.StatusCode, "Should return 200 for sucessfull")
	assert.Equal(s.T(), user.UserName, s.user.UserName, "should have the same username of previously created")
}

func TestGetUserEndpointSuite(t *testing.T) {
	suite.Run(t, new(GetUserTestSuite))
}
