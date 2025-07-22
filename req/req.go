package req

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// MakeReq envia uma requisição HTTP com método, rota, corpo opcional e token de autenticação.
// Retorna o corpo da resposta em bytes ou um erro.
//
// Parâmetros:
//   - method: "GET", "POST", etc.
//   - baseUrl: URL base da API
//   - route: caminho adicional da requisição (ex: "/Assistance/Ticket")
//   - body: ponteiro para string JSON (ou nil)
//   - auth: token de autenticação
func MakeReq(method string, baseUrl string, route string, body *string, auth string) ([]byte, error) {
	route = baseUrl + route
	var reqBody io.Reader

	if body != nil {
		reqBody = strings.NewReader(*body)
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(method, route, reqBody)
	if err != nil {
		return []byte(err.Error()), err
	}

	req.Header.Set("Authorization", auth)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte(err.Error()), err
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	return response, nil
}
