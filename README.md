# GLPI Go Wrapper

Um wrapper simples em Go para a API v2 do GLPI, compatível com a versão 11.0.0+ que será lançada.

## ⚠️ Aviso Importante

Este é um dos meus primeiros projetos em Go, então podem ocorrer erros ou bugs. Use com cautela em ambiente de produção e sinta-se à vontade para contribuir com melhorias!

## 📋 Pré-requisitos

- Go 1.24.4 ou superior
- GLPI versão 11.0.0+ (quando lançada)
- Credenciais OAuth2 configuradas no GLPI

## 🚀 Instalação

```bash
git clone <seu-repositorio>
cd gtk
go mod tidy
```

## ⚙️ Configuração

Antes de usar, configure as constantes no arquivo `main.go`:

```go
const API_URL = "http://seu-glpi.com/api.php/"
const USERNAME = "seu-usuario"
const PASSWORD = "sua-senha"
const CLIENT_ID = "seu-client-id"
const CLIENT_SECRET = "seu-client-secret"
```

## 📖 Uso Básico

### Autenticação

```go
import (
    authMethods "gtk/auth"
    "gtk/types"
)

// Obter token de autenticação
authBytes, err := authMethods.GetAuthenticated(
    API_URL, 
    USERNAME, 
    PASSWORD, 
    CLIENT_ID, 
    CLIENT_SECRET
)
if err != nil {
    log.Panic(err)
}

var auth types.AuthParams
json.Unmarshal(authBytes, &auth)
```

### Fazendo Requisições

```go
import "gtk/req"

// Fazer uma requisição GET
response, err := req.MakeReq(
    "GET", 
    API_URL, 
    "/Assistance/Ticket", 
    nil, 
    auth.Token
)
if err != nil {
    log.Panic(err)
}

fmt.Println(string(response))
```

## 📁 Estrutura do Projeto

```
gtk/
├── auth/
│   └── auth.go          # Funções de autenticação OAuth2
├── req/
│   └── req.go           # Funções para requisições HTTP
├── types/
│   ├── auth.go          # Struct para dados de autenticação
│   └── authParams.go    # Struct para resposta de autenticação
├── main.go              # Exemplo de uso
├── go.mod               # Dependências do módulo
└── README.md            # Este arquivo
```

## 🔧 Funcionalidades Implementadas

- ✅ Autenticação OAuth2 (Password Grant)
- ✅ Requisições GET/POST/PUT/DELETE
- ✅ Gerenciamento automático de headers
- ✅ Suporte a tokens de acesso

## 🚧 Limitações Conhecidas

- Tratamento básico de erros
- Sem retry automático para falhas de rede
- Sem cache de tokens
- Sem validação de SSL configurável

## 🤝 Contribuindo

Como este é um projeto de aprendizado, contribuições são muito bem-vindas! Sinta-se à vontade para:

1. Reportar bugs
2. Sugerir melhorias
3. Enviar pull requests
4. Compartilhar feedback

## 📝 Changelog

### v0.1.0
- Implementação básica de autenticação OAuth2
- Suporte a requisições HTTP básicas
- Estrutura inicial do projeto

---

**Nota**: Este wrapper foi desenvolvido com base na documentação prévia da API v2 do GLPI. Podem ser necessários ajustes quando a versão 11.0.0 for oficialmente lançada.
