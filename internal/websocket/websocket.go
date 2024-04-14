package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao fazer upgrade da conexão HTTP para WebSocket:", err)
		return
	}
	defer conn.Close()

	// Loop para ler mensagens do cliente e escrever de volta
	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler mensagem do cliente:", err)
			continue // Se houver um erro na leitura da mensagem, interrompa o loop, mas mantenha a conexão aberta
		}
		conn.WriteMessage(websocket.TextMessage, []byte(os.Getenv("APPID")))
		// Decodificar a mensagem JSON do cliente em uma estrutura de dados
		var data map[string]string
		err = json.Unmarshal(message, &data)
		if err != nil {
			log.Println("Erro ao decodificar JSON:", err)
			continue // Se houver um erro na decodificação do JSON, ignore esta mensagem e continue com a próxima
		}
		log.Printf("Received message:%s %s", os.Getenv("APPID"), message)

		// Extrair as informações de nome e tagline do payload WebSocket
		username := data["username"]
		tagline := data["tagline"]

		// Construir a URL formatada
		url := "https://api.antonyz.tk/lol/rank?name=" + username + "&tag=" + tagline + "&r=BR&lang=pt"

		// Criar uma nova instância de resty
		client := resty.New()

		// Enviar a solicitação GET para a URL formatada
		resp, err := client.R().
			EnableTrace().
			Get(url)
		if err != nil {
			log.Println("Erro ao fazer solicitação GET:", err)
			continue // Se houver um erro na solicitação, interrompa o loop, mas mantenha a conexão aberta
		}

		// Verificar se a resposta indica "NAO RANQUEADO"
		bodyString := string(resp.Body())
		if bodyString == "NAO RANQUEADO" {
			conn.WriteMessage(websocket.TextMessage, resp.Body())
			continue // Se o jogador não estiver ranqueado, envie a resposta e continue com a próxima mensagem
		}

		// Escrever a resposta de volta para o cliente WebSocket
		ranking, division, pdl := extractRankingDivisionLP(bodyString)

		// Formatar os dados extraídos em uma estrutura
		data = map[string]string{
			"ranking":  ranking,
			"division": division,
			"pdl":      pdl,
		}

		// Serializar a estrutura para JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Erro ao serializar dados para JSON:", err)
			continue // Se houver um erro na serialização para JSON, interrompa o loop, mas mantenha a conexão aberta
		}

		// Escrever a resposta de volta para o cliente WebSocket
		err = conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Println("Erro ao escrever mensagem de volta para o cliente:", err)
			continue // Se houver um erro ao escrever a mensagem, interrompa o loop, mas mantenha a conexão aberta
		}
	}
}

func extractRankingDivisionLP(response string) (string, string, string) {
	// Use expressões regulares para extrair as informações
	re := regexp.MustCompile(`(.+) \((\d+) PDL\)`)
	matches := re.FindStringSubmatch(response)
	if len(matches) < 3 {
		return "", "", ""
	}

	// O primeiro grupo de captura contém o ranking
	ranking := matches[1]

	// O segundo grupo de captura contém a divisão
	division := matches[2]

	// O terceiro grupo de captura contém os LPs
	pdl := matches[3]

	return ranking, division, pdl
}
