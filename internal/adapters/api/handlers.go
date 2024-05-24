package api

import (
	"encoding/json"
	"go-bitcoin-ltp/internal/core/services"
	"net/http"
)

type Handler struct {
	ltpService *services.LTPService
}

func NewHandler(service *services.LTPService) *Handler {
	return &Handler{ltpService: service}
}

func (h *Handler) GetLTP(w http.ResponseWriter, r *http.Request) {
	ltp, err := h.ltpService.GetLastTradedPrices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"ltp": ltp,
	}

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode("")  // This returns additional \n at the end - cool feature
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
