# Learning Labs Auction - GoExpert

Esta aplicação gerencia leilões (auctions) e lances (bids) em Go, utilizando MongoDB.

## Execução

1. **Subir os containers Docker**:

   ```bash
   docker-compose up -d
   ```

2. A aplicação estará disponível em `localhost:8080`.

## Testando Endpoints

Use a extensão REST Client do VSCode com os arquivos em `api/`:

- **auctions.http**  
  - `POST /auction`: Cria um leilão  
  - `GET /auction/{id}`: Retorna um leilão  
  - `GET /auction?status={status}`: Retorna leilões por status  
  - `GET /auction/winner/{id}`: Retorna o lance vencedor

- **bids.http**  
  - `POST /bid`: Cria um lance  
  - `GET /bid/{auctionId}`: Retorna lances de um leilão

- **users.http**  
  - `GET /user/{id}`: Retorna dados de um usuário

## Testes

### Pré-requisitos

- Docker e Docker Compose  
- Go 1.20+

### Iniciando MongoDB para Testes

1. **Subir apenas o MongoDB**:

   ```bash
   docker-compose up -d mongodb
   ```

2. **Verificar se está rodando**:

   ```bash
   docker ps
   ```

### Executando Testes

```bash
go test ./internal/infra/database/auction/... -v
```

Saída de exemplo:

```go
❯ go test ./internal/infra/database/auction/... -v
=== RUN   TestCreateAuctionAndUpdateStatus
--- PASS: TestCreateAuctionAndUpdateStatus (3.01s)
=== RUN   TestCreate10AuctionsAndUpdateStatus
--- PASS: TestCreate10AuctionsAndUpdateStatus (3.02s)
PASS
ok      fullcycle-auction_go/internal/infra/database/auction    6.032s
```

### Finalizando

```bash
docker-compose down
```
