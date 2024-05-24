package main

import (
	"go-bitcoin-ltp/internal/adapters/api"
	"go-bitcoin-ltp/internal/core/services"
	"go-bitcoin-ltp/pkg/testutils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLTP(t *testing.T) {
	mockClient := &testutils.MockAPIClient{}
	service := services.NewLTPService(mockClient)
	handler := api.NewHandler(service)

	req, err := http.NewRequest("GET", "/api/v1/ltp", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.GetLTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ltp":[{"pair":"BTC/CHF","amount":"49000.12"},{"pair":"BTC/EUR","amount":"50000.12"},{"pair":"BTC/USD","amount":"52000.12"}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
