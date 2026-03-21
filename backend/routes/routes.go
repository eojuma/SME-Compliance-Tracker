package routes

import (
	"net/http"

	"SME-compliance-tracker/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/api/v1", handlers.APIRoot)

	return mux
}