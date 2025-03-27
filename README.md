# Learning Labs Auction - GoExpert

## Como Executar os Testes

### Pré-requisitos

- Docker e Docker Compose instalados

- Go 1.20+ instalado

### Iniciando o MongoDB para Testes

1. Suba apenas o container do MongoDB:

```bash
docker-compose up -d mongodb
```

1. Verifique se o MongoDB está rodando:

```bash
docker ps
```

### Executando os Testes

Para rodar todos os testes de integração com o MongoDB:

```bash
go test ./internal/infra/database/auction/... -v
```

### Exemplo de Teste de Integração

Os testes de integração verificam:

- Conexão com o banco de dados

- Atualização automática de status

Exemplo de saída esperada:

```go
=== RUN   TestCreateAuctionAndUpdateStatus
--- PASS: TestCreateAuctionAndUpdateStatus (3.01s)
=== RUN   TestCreate10AuctionsAndUpdateStatus
--- PASS: TestCreate10AuctionsAndUpdateStatus (3.02s)
PASS
ok      fullcycle-auction_go/internal/infra/database/auction    6.032s
```

### Parando o MongoDB

Após os testes, você pode parar o container com:

```bash
docker-compose down
```
