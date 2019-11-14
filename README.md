# Crawler

Crawler para buscar o valor atual de uma Transferência do Plano Profissional no site da SmartMei.

## Dependências

 - [Golang 1.13](https://golang.org/) - Linguagem usada para implementar a solução.
 - [gqlgen](https://github.com/99designs/gqlgen/) - Implementação servidor GraphQL.

## Execução

Para executar o projeto use o seguinte comando:

```sh
go run server/server.go -url=https://www.smartmei.com.br
# 2019/11/13 20:04:12 connect to http://localhost:8080/ for GraphQL playground
```

 - Parâmetro url - Endereço do site da SmartMei 

### Playground

```sh
http://localhost:8080/
```

#### Consulta

Exemplo de Consulta

```sh
query {
  tarifaAtual {
    data
    descricao
    valorReal
    valorEuro
    valorDolar
  }
}
```

Saída

```json
{
  "data": {
    "tarifaAtual": {
      "data": "2019-11-13 20:08:26",
      "descricao": "R$ 7,00",
      "valorReal": 7,
      "valorEuro": 1.5211769566,
      "valorDolar": 1.6742073578999999
    }
  }
}
```

## Testes

Para executar os testes do projeto use o seguinte comando:

```sh
go test ./...

# ?       bitbucket.org/challen_ge/smartmei_backend       [no test files]
# ok      bitbucket.org/challen_ge/smartmei_backend/internal/cotacao      0.003s
# ok      bitbucket.org/challen_ge/smartmei_backend/internal/crawler      0.628s
# ?       bitbucket.org/challen_ge/smartmei_backend/server        [no test files]
```