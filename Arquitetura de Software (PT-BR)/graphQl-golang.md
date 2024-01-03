
Para criar um exemplo simples de servidor GraphQL em Go, você precisará de uma biblioteca GraphQL para Go. Uma das bibliotecas mais populares é a graphql-go/graphql. Este exemplo abordará a criação de um servidor GraphQL simples com um único tipo e uma query.

## Passos:

### Instalar as Dependências:
Você precisará instalar a biblioteca graphql-go com o seguinte comando:
```
go get github.com/graphql-go/graphql
```

### Escrever o Código:

Aqui está um exemplo simples de um servidor GraphQL que responde a uma query para buscar um "Hello" message.

main.go
```
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/graphql-go/graphql"
)

func main() {
    // Schema
    fields := graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "world", nil
            },
        },
    }

    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

    // Query
    query := `
        {
            hello
        }
    `
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }

    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}
```

Este código define um esquema GraphQL com uma única query chamada "hello" que retorna a string "world". Quando você executa este programa, ele imprimirá o resultado da query GraphQL, que é a resposta à query "hello".

### Executar o Código:

- Salve o código em um arquivo chamado main.go.
- Execute o código com o comando go run main.go.

Quando você executar este programa, ele imprimirá a resposta da query GraphQL no console, que será algo parecido com:

json
```
{"data":{"hello":"world"}}
```

### Notas:

- Este é um exemplo muito básico para demonstrar uma implementação simples do GraphQL em Go.
- Para uma aplicação real, você provavelmente irá querer aceitar requests HTTP, lidar com queries variadas e talvez conectar-se a um banco de dados ou outras APIs para buscar e manipular dados.
- A biblioteca graphql-go tem muitos outros recursos que você pode explorar para criar esquemas mais complexos, mutações, subscriptions e mais.
