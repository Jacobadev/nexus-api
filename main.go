package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port uint16 = 1337

type Endereco struct {
	Id     int    `json:"id"`
	Rua    string `json:"titulo"`
	Numero string `json:"numero"`
	CEP    string `json:"cep"`
}

var Enderecos []Endereco = []Endereco{
	{
		Id:     1,
		Rua:    "Guilherme",
		Numero: "654",
		CEP:    "24422330",
	},
	{
		Id:     2,
		Rua:    "Alzira",
		Numero: "123",
		CEP:    "24422320",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bem vindo :)</h1>")
}

//GET enderecos
func buscarEndereco(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Enderecos)
}

func rotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/address", buscarEndereco)
}

func main() {
	rotas()

	fmt.Printf("Servidor rodando na %v \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil)) //DefaultServerMux
}
