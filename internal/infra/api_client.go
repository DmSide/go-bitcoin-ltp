package infra

import (
	"encoding/json"
	"fmt"
	"go-bitcoin-ltp/internal/domain"
	"io"
	"net/http"
)

type APIClient interface {
	FetchLTP(pair string) (domain.LTP, error)
}

type apiClient struct{}

func NewAPIClient() APIClient {
	return &apiClient{}
}

func (c *apiClient) FetchLTP(pair string) (domain.LTP, error) {
	const UrlTemplate = "https://api.kraken.com/0/public/Ticker?pair="
	url := UrlTemplate + pair

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return domain.LTP{}, fmt.Errorf("error fetching data for %s: %w", pair, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.LTP{}, fmt.Errorf("error reading response body for %s: %w", pair, err)
	}

	var response domain.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return domain.LTP{}, fmt.Errorf("error unmarshalling JSON for %s: %w", pair, err)
	}

	for key, value := range response.Result {
		return domain.LTP{Pair: key, Amount: value.C[0]}, nil
	}

	return domain.LTP{}, fmt.Errorf("no data found for %s", pair)
}
