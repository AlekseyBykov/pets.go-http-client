# CoinCap HTTP Client (Go)

A lightweight Go client for interacting with the public CoinCap API (https://api.coincap.io). It retrieves cryptocurrency asset data.

## Description

This application performs HTTP requests to the CoinCap API and displays information about:
- All available assets (`/v2/assets`)
- A specific asset by ID (`/v2/assets/{id}`)

Includes a custom `loggingRoundTripper` for HTTP request logging.

## Usage

### Run the application
```bash
go run ./cmd/main.go
```
### Sample output
```bash
[ID] bitcoin | [RANK] 1 | [SYMBOL] BTC | [NAME] Bitcoin | [PRICE] 65000.00
[ID] ethereum | [RANK] 2 | [SYMBOL] ETH | [NAME] Ethereum | [PRICE] 3500.00
...
[ID] bitcoin | [RANK] 1 | [SYMBOL] BTC | [NAME] Bitcoin | [PRICE] 65000.00
```
## Features

- HTTP client with timeout and request logging.
- Parses JSON responses from CoinCap API.
- Simple and extendable structure for future endpoints.

## Dependencies
Go 1.20+
Standard Go library (net/http, encoding/json, etc.)


