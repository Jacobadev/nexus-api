package http

import (
	"net/http"
	"testing"

	"github.com/gateway-address/pkg/httpErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	URL string
}

func (s *ExampleTestSuite) SetupTest() {
	s.URL = "http://localhost:5000/api/v1/auth/"
}

func (s *ExampleTestSuite) TestShouldReturnOK() {
	c := http.Client{}
	res, err := c.Post(s.URL, "", nil)
	if err != nil {
		s.T().Fatal(err)
	}
	defer res.Body.Close()
	assert.Equal(s.T(), http.StatusCreated, res.StatusCode)
}

func (s *ExampleTestSuite) TestShouldReturnBadRequest() {
	c := http.Client{}
	res, err := c.Post(s.URL, "", nil)
	if err != nil {
		s.T().Fatal(err)
	}
	defer res.Body.Close()
	assert.Equal(s.T(), http.StatusBadRequest, httpErrors.BadRequest)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
