package services

import (
	"go-bitcoin-ltp/internal/domain"
	"go-bitcoin-ltp/internal/infra"
	"sync"
)

const MaxCacheSize = 3

type LTPService struct {
	apiClient infra.APIClient
	cache     map[string]domain.LTP
	cacheKeys [MaxCacheSize]string
	cacheLock sync.Mutex
}

func NewLTPService(client infra.APIClient) *LTPService {
	return &LTPService{
		apiClient: client,
		cache:     make(map[string]domain.LTP, MaxCacheSize),
		cacheKeys: [MaxCacheSize]string{"BTC/CHF", "BTC/EUR", "BTC/USD"},
	}
}

func (s *LTPService) GetLastTradedPrices() ([]domain.LTP, error) {
	currencyPairs := s.cacheKeys
	var ltpList []domain.LTP

	for _, pair := range currencyPairs {
		ltp, err := s.getLTP(pair)
		if err != nil {
			return nil, err
		}
		ltpList = append(ltpList, ltp)
	}

	return ltpList, nil
}

func (s *LTPService) getLTP(pair string) (domain.LTP, error) {
	s.cacheLock.Lock()
	defer s.cacheLock.Unlock()

	ltp, err := s.apiClient.FetchLTP(pair)
	if err != nil {
		if cachedValue, ok := s.cache[pair]; ok {
			return cachedValue, nil
		}
		return domain.LTP{Pair: pair, Amount: "0"}, nil
	}

	s.cache[pair] = ltp

	return ltp, nil
}
