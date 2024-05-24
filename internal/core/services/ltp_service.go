package services

import (
	"go-bitcoin-ltp/internal/domain"
	"go-bitcoin-ltp/internal/infra"
)

type LTPService struct {
	apiClient infra.APIClient
}

func NewLTPService(client infra.APIClient) *LTPService {
	return &LTPService{apiClient: client}
}

func (s *LTPService) GetLastTradedPrices() ([]domain.LTP, error) {
	return s.apiClient.FetchLTP()
}
