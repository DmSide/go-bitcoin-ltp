# Go Bitcoin LTP Service

This is a service written in Go that provides an API for retrieval of the Last Traded Price (LTP) of Bitcoin for the following currency pairs:
- BTC/USD
- BTC/CHF
- BTC/EUR

## Requirements

- Docker
- Go (for local development)

## Building and Running the Application

### Using Docker

1. Build the Docker image:
   ```sh
   docker build -t go-bitcoin-ltp .
   ```

2. Run the Docker container:
   ```sh
   docker run -p 8080:8080 go-bitcoin-ltp
   ```

The server will start on port 8080.

### Running Locally

1. Install dependencies:
   ```sh
   go mod download
   ```

2. Build the application:
   ```sh
   cd cmd/server
   go build -o ../../go-bitcoin-ltp
   cd ../..
   ```

3. Run the application:
   ```sh
   ./go-bitcoin-ltp
   ```

The server will start on port 8080.

## Testing

Run 
 ```sh
   go test ./...
```


## API Endpoint
### GET /api/v1/ltp

  Response:
  ```json
  {
    "ltp": [
      {
        "pair": "BTC/CHF",
        "amount": "49000.12"
      },
      {
        "pair": "BTC/EUR",
        "amount": "50000.12"
      },
      {
        "pair": "BTC/USD",
        "amount": "52000.12"
      }
    ]
  }
