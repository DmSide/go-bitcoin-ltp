package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go-bitcoin-ltp/internal/adapters/api"
	"go-bitcoin-ltp/internal/core/services"
	"go-bitcoin-ltp/internal/infra"
)

func main() {
	apiClient := infra.NewAPIClient()
	ltpService := services.NewLTPService(apiClient)
	apiHandler := api.NewHandler(ltpService)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/ltp", apiHandler.GetLTP).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
