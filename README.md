# Trading System API Setup

## Getting Started

### Prerequisites
- Go 1.20+
- Docker

### Installation

1. Clone the repository:
   ```bash
git clone https://github.com/schen7239/autonomous-agent-repo-test.git
```
2. Navigate to the project directory:
   ```bash
cd autonomous-agent-repo-test
```
3. Install dependencies:
   ```bash
go mod tidy
```

### Running the API

1. Start the API server:
   ```bash
go run main.go
```

### Using the Endpoints

#### CreateBuy
- Endpoint: `/createbuy`
- Method: `POST`
- Request Body: `{"symbol": "AAPL", "quantity": 10}`

#### CreateSell
- Endpoint: `/createsell`
- Method: `POST`
- Request Body: `{"symbol": "AAPL", "quantity": 10}`

#### CommitTrade
- Endpoint: `/committrade`
- Method: `POST`
- Request Body: `{"tradeId": "12345"}`

## Contributing

1. Fork the repository.
2. Create a new branch for your feature/fix.
3. Commit your changes.
4. Push to your branch.
5. Create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

