 Exemplo para uma autenticação de API Gateway em Go, com uma única camada de segurança — utilizando JWT para autenticação e TLS para comunicação segura. Este exemplo assumirá uso de HTTPS (TLS) para todas as comunicações.

 ### Requisitos:

- Go instalado em seu sistema.
- Certificados TLS para HTTPS.
- Biblioteca de JWT para Go (por exemplo, github.com/dgrijalva/jwt-go).

### API GATEWAY:

O Gateway será responsável por autenticar as requisições usando JWT e encaminhá-las para os serviços internos.

```
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

// Secret key used to sign tokens (em produção, use uma variável de ambiente ou configuração segura)
var jwtKey = []byte("my_secret_key")

// Handler da API que atua como um gateway
func apiHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    // Parse e valida o token JWT
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintf(w, "Unauthorized: %s", err.Error())
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Aqui você pode extrair informações do token e adicionar lógica para roteamento
        log.Printf("User %v accessed", claims["user"])
        // Encaminhar a solicitação para o serviço interno ou responder diretamente
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintln(w, "Unauthorized")
    }
}

func main() {
    http.HandleFunc("/api", apiHandler)

    // Certifique-se de usar HTTPS em produção
    log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}
```

Neste código:

- Usamos o pacote jwt-go para decodificar e validar o token JWT.
- O Gateway escuta na porta 8080 e espera que as requisições incluam um token JWT válido.
- Este exemplo usa HTTPS (TLS), então você precisa de um certificado server.crt e uma chave server.key.

### Observações:
- Este exemplo é simplificado para fins de demonstração. Em um cenário real, você precisará gerenciar tokens, chaves e certificados com segurança.
- O manuseio de erros é fundamental, especialmente em produção.
- A comunicação entre o Gateway e os serviços internos também deve ser segura e autenticada.
- Em um sistema real, você também precisará lidar com a renovação de tokens, logout, e outros aspectos de gestão de sessão e segurança.

Essa simplificação foca em uma única camada de segurança baseada em JWT e TLS/SSL para o transporte.
