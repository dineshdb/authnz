package api

import (
	"net/http"

	"github.com/dineshdb/authnz/internal/utils"
)

// Liveness and readiness probes for kubernetes deployments. Health check are important part of cloud native infrastructure
func Live(w http.ResponseWriter, r *http.Request) {
	utils.OK(w, nil)
}

func Ready(w http.ResponseWriter, r *http.Request) {
	// TODO: Add readiness checks for dependencies once you add external services like postgres
	utils.OK(w, nil)
}
