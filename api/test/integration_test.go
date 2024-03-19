package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gateway-address/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestEndToEnd(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestHappyHealthCheck() {
	c := http.Client{}
	r, err := c.Get("http://:3333/api/v1/health")
	require.Nil(s.T(), err, "Error fetching health check response") // Assertions with error handling

	defer r.Body.Close() // Ensure body is closed even on error

	b, err := io.ReadAll(r.Body)
	require.Nil(s.T(), err, "Error reading response body")

	// Expected response
	expectedResponse := `{"status": "OK"}`

	// Compare JSON using testify's JSONEq assertion
	require.JSONEq(s.T(), string(b), expectedResponse)
}

func (s *TestSuite) TestCreateUserResponse() {
	body := []byte(`{
		"first_name": "John",
		"last_name": "Doe",
		"username": "johndoe",
		"email": "johndoe@example.com",
		"password": "password123"
	}`)
	c := http.Client{}
	r, err := c.Post("http://0.0.0.0:3333/api/v1/user", "application/json", bytes.NewBuffer(body))
	require.Nil(s.T(), err, "Error fetching health check response") // Assertions with error handling

	require.Nil(s.T(), err, "Error reading response body")
	assert.Equal(s.T(), http.StatusCreated, r.StatusCode)
}

func (s *TestSuite) TestCreateUserResponseWithInvalidBody() {
	// Corpo JSON com campo faltando
	invalidBody := []byte(`{
        "first_name": "John",
        "last_name": "Doe",
        "email": "johndoe@example.com",
        "password": "password123"
    }`)

	// Envia uma requisição POST com o corpo JSON incorreto
	c := http.Client{}
	r, err := c.Post("http://0.0.0.0:3333/api/v1/user", "application/json", bytes.NewBuffer(invalidBody))
	require.Nil(s.T(), err, "Error sending POST request")

	// Verifica se a resposta tem o status HTTP 400 Bad Request
	assert.Equal(s.T(), http.StatusBadRequest, r.StatusCode)
}

func (s *TestSuite) TestGetUsersResponse() {
	// Cria um cliente HTTP
	c := http.Client{}

	r, err := c.Get("http://localhost:3333/api/v1/user")
	require.Nil(s.T(), err, "Error fetching users response")

	assert.Equal(s.T(), http.StatusOK, r.StatusCode)

	// Lê o corpo da resposta
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	require.Nil(s.T(), err, "Error reading response body")

	// Verifica se os dados dos usuários estão presentes no corpo da resposta
	// Implemente suas próprias verificações conforme necessário
	// Exemplo: verifica se o corpo da resposta contém um array de usuários
	var users []user.User
	err = json.Unmarshal(body, &users)
	require.Nil(s.T(), err, "Error unmarshalling response body")

	// Agora você pode fazer as asserções necessárias nos dados dos usuários
	// Por exemplo, você pode verificar se a lista de usuários não está vazia
	assert.NotEmpty(s.T(), users, "No users returned in the response")
}
