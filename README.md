# Trading System API

## Setup Instructions

1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Run `go run src/main.go` to start the API server.
4. Use endpoints `/createBuy`, `/createSell`, and `/commitTrade` to manage trades.

## Endpoints

- `/createBuy`: Create a buy order.
- `/createSell`: Create a sell order.
- `/commitTrade?id=<id>`: Commit a trade by ID.

## Example

```bash
# Create a buy order
curl http://localhost:8080/createBuy

# Create a sell order
curl http://localhost:8080/createSell

# Commit a trade
curl http://localhost:8080/commitTrade?id=<id>
```

