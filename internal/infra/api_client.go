package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-bitcoin-ltp/internal/domain"
	_ "net/http"
	_ "time"
)

type APIClient interface {
	FetchLTP() ([]domain.LTP, error)
}

type apiClient struct{}

func NewAPIClient() APIClient {
	return &apiClient{}
}

func (c *apiClient) FetchLTP() ([]domain.LTP, error) {
	// Mock response for example purposes
	mockResponse := `[
        {"pair": "BTC/CHF", "amount": "49000.12"},
        {"pair": "BTC/EUR", "amount": "50000.12"},
        {"pair": "BTC/USD", "amount": "52000.12"}
    ]`

	var ltp []domain.LTP
	err := json.Unmarshal([]byte(mockResponse), &ltp)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to parse LTP data %s", err.Error()))
	}

	return ltp, nil
}
