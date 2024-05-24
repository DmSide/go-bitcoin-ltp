package testutils

import (
	"go-bitcoin-ltp/internal/domain"
)

type MockAPIClient struct{}

func (m *MockAPIClient) FetchLTP() ([]domain.LTP, error) {
	return []domain.LTP{
		{Pair: "BTC/CHF", Amount: "49000.12"},
		{Pair: "BTC/EUR", Amount: "50000.12"},
		{Pair: "BTC/USD", Amount: "52000.12"},
	}, nil
}
