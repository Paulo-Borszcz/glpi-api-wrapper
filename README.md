<div align="center">

# GLPI Go Wrapper

### _Um wrapper em Go para a API v2 do GLPI_

[![Go Version](https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![GLPI](https://img.shields.io/badge/GLPI-11.0.0+-FF6B35?style=for-the-badge&logo=glpi)](https://glpi-project.org/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow?style=for-the-badge)]()

</div>

---

## Sobre o Projeto

Este é um wrapper simples e intuitivo para a **API v2 do GLPI**, desenvolvido especificamente para ser compatível com a versão **11.0.0+** que será lançada em breve.

> **🌱 Projeto de Aprendizado**  
> Este é um dos meus primeiros projetos em Go! Embora funcional, pode conter bugs ou áreas para melhoria. Contribuições e feedback são super bem-vindos! 💚

## Pré-requisitos

<table>
<tr>
<td>

**Tecnologias**
- Go 1.24.4+
- GLPI 11.0.0+

</td>
<td>

**Configurações**
- Credenciais OAuth2 no GLPI
- Cliente OAuth2 configurado

</td>
</tr>
</table>

## Instalação Rápida

```bash
# Clone o repositório
git clone https://github.com/Paulo-Borszcz/glpi-api-wrapper.git

# Entre no diretório
cd gtk

# Baixe as dependências
go mod tidy
```

> **Dica**: Certifique-se de ter o Go instalado em sua máquina antes de prosseguir!

## Configuração

<details>
<summary><b>Configurar Credenciais (Clique para expandir)</b></summary>

<br>

Edite o arquivo `main.go` com suas credenciais:

```go
const API_URL = "http://seu-glpi.com/api.php/"      // URL do seu GLPI
const USERNAME = "seu-usuario"                       // Seu usuário
const PASSWORD = "sua-senha"                         // Sua senha
const CLIENT_ID = "seu-client-id"                    // ID do cliente OAuth2
const CLIENT_SECRET = "seu-client-secret"            // Secret do cliente OAuth2
```

</details>

## Exemplos de Uso

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

// Fazer uma requisição GET para tickets
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

fmt.Println("Resposta:", string(response))
```

## Arquitetura do Projeto

```
🗂️ gtk/
├──  auth/
│   └── auth.go              # Magia da autenticação OAuth2
├──  req/
│   └── req.go               # Motor das requisições HTTP  
├──  types/
│   ├── auth.go              # Estrutura dos dados de auth
│   └── authParams.go        # Estrutura da resposta de auth
├──  main.go                # Exemplo prático de uso
├──  go.mod                # Dependências do projeto
└──  README.md             # Você está aqui! 👋
```

## Funcionalidades

<div align="center">

| Funcionalidade | Status | Descrição |
|---|---|---|
| **OAuth2** | ✅ | Autenticação Password Grant |
| **HTTP Methods** | ✅ | GET/POST/PUT/DELETE |
| **Headers** | ✅ | Gerenciamento automático |
| **Tokens** | ✅ | Suporte completo |

</div>

## Conhecendo as Limitações

<details>
<summary><b>🔍 O que ainda pode melhorar (Clique para ver)</b></summary>

<br>

- **Retry automático** para falhas de rede
- **Cache de tokens** para melhor performance  
- **Validação SSL** configurável
- **Logs** mais detalhados
- **Rate limiting** inteligente

> Essas são oportunidades perfeitas para contribuir! 🤝

</details>

## 🤝 Como Contribuir

<div align="center">

### Sua contribuição faz a diferença! 💚

</div>

Este é um projeto de **aprendizado colaborativo**. Toda ajuda é bem-vinda:

<table>
<tr>
<td align="center">

**🐛 Encontrou um Bug?**  
[Reporte aqui](../../issues)

</td>
<td align="center">

**💡 Tem uma Ideia?**  
[Sugira aqui](../../issues)

</td>
<td align="center">

**🔧 Quer Codificar?**  
[Pull Request](../../pulls)

</td>
</tr>
</table>

### 📋 Como Contribuir:

1. **Fork** o projeto
2. **Crie** uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. **Commit** suas mudanças (`git commit -m 'Add: Minha nova feature'`)
4. **Push** para a branch (`git push origin feature/MinhaFeature`)
5. **Abra** um Pull Request

> **💭 Dica**: Mesmo pequenas correções de typos são valiosas!

## 📚 Links Úteis

<div align="center">

[![GLPI Official](https://img.shields.io/badge/📖_GLPI-Official_Docs-FF6B35?style=for-the-badge&logo=glpi)](https://glpi-project.org/)
[![Go Documentation](https://img.shields.io/badge/📚_Go-Documentation-00ADD8?style=for-the-badge&logo=go)](https://golang.org/doc/)
[![API v2](https://img.shields.io/badge/🔗_API-v2_Reference-blue?style=for-the-badge)](https://github.com/glpi-project/glpi)

</div>

## Histórico de Versões

<details>
<summary><b>🏷️ Releases (Clique para expandir)</b></summary>

<br>

### 🎉 v0.1.0 - _Primeira Release_
- Implementação básica de autenticação OAuth2
- Suporte a requisições HTTP básicas  
- Estrutura inicial do projeto
- Documentação inicial

### Próximas Versões
- Sistema de retry automático
- Cache inteligente de tokens
- Logging avançado
- Testes automatizados

</details>

---

<div align="center">

### 🌟 Se este projeto te ajudou, deixe uma ⭐!

**Desenvolvido com ❤️ e muito ☕**  
_"Todo expert já foi um iniciante"_

[![Made with Love](https://img.shields.io/badge/Made%20with-❤️-red?style=for-the-badge)]()
[![Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?style=for-the-badge&logo=go)]()

</div>
