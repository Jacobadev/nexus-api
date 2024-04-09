package e2etest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/auth/repository"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/db/postgres"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	URL string
}

func (s *LoginTestSuite) setupLoginTest() {
	s.URL = "http://localhost:5000/api/v1/auth/login"

	cfgFile, err := config.LoadConfig("../../../config/config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	// Inicializar e conectar ao banco de dados
	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("NewPsqlDB: %v", err)
	}
	repo := repository.NewRepositorySqlite(db)
	if existsUser, err := repo.FindByEmail(LoginDetails); existsUser != nil || err == nil {
		return
	}

	defer db.Close()

	// Gerar token JWT para o usuário de login
	err = LoginDetails.HashPassword()
	if err != nil {
		s.T().Errorf("HashPassword: %v", err)
	}

	_, err = db.Exec("INSERT INTO users (first_name, last_name, email, username, password) VALUES ($1, $2, $3, $4, $5)",
		LoginDetails.FirstName, LoginDetails.LastName, LoginDetails.Email, LoginDetails.UserName, LoginDetails.Password)
	if err != nil {
		s.T().Errorf("DB Exec: %v", err)
	}
}

// LoginDetails contém os detalhes do login

var LoginDetails = &model.User{
	FirstName: "Reine",
	LastName:  "Freshwater",
	Email:     "rfreshwater1@godaddy.com",
	UserName:  "rfreshwater1",
	Password:  "rxA!mp_password",
}

var LoginDetailsNoHash = &model.User{
	FirstName: "Reine",
	LastName:  "Freshwater",
	Email:     "rfreshwater1@godaddy.com",
	UserName:  "rfreshwater1",
	Password:  "rxA!mp_password",
}

func (s *LoginTestSuite) TestLoginShouldReturnOK() {
	s.setupLoginTest()
	loginJSON, err := json.Marshal(LoginDetailsNoHash)
	if err != nil {
		s.T().Errorf("Erro ao transformar os detalhes do login em JSON: %v", err)
		return
	}

	// Crie um cliente HTTP
	client := http.Client{}

	// Envie a solicitação POST com os detalhes de login formatados como JSON
	res, err := client.Post(s.URL, "application/json", bytes.NewBuffer(loginJSON))
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
	assert.Equal(s.T(), LoginDetails.FirstName, responseBody.User.FirstName)
	assert.Equal(s.T(), LoginDetails.LastName, responseBody.User.LastName)
	assert.Equal(s.T(), LoginDetails.Email, responseBody.User.Email)
	assert.Equal(s.T(), LoginDetails.UserName, responseBody.User.UserName)
}

func TestLoginEndpointSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
