package auth

import (
	"encoding/json"
	"fmt"
	"gtk/types"
	"io"
	"net/http"
	"strings"
)

// GetAuthenticated realiza uma requisição de autenticação do tipo "password grant".
// Envia as credenciais do usuário e dados do client para obter um token de acesso.
//
// Parâmetros:
//   - base_url: URL base da API (ex: "https://glpi.com/api.php/")
//   - username: nome de usuário para autenticação
//   - password: senha do usuário
//   - clientId: identificador do cliente OAuth2
//   - clientSecret: segredo do cliente OAuth2
//
// Retorna:
//   - []byte: corpo da resposta em bytes (espera-se um JSON com o token)
func GetAuthenticated(base_url string, username string, password string, clientId string, clientSecret string) ([]byte, error) {
	reqBody := types.Auth{
		GrantType:    "password",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		Scope:        "api",
	}

	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return []byte(err.Error()), err
	}

	reqUrl := fmt.Sprint(base_url, "token")
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(string(reqBodyJson)))
	if err != nil {
		return []byte(err.Error()), err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte(err.Error()), err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte(err.Error()), err
	}

	return body, nil
}
