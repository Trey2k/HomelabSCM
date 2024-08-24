package api

import (
	"encoding/json"
	"net/http"

	"homelabscm.com/scm/internal/pkg/config"
	"homelabscm.com/scm/internal/pkg/logger"
	"homelabscm.com/scm/pkg/api_model"
)

func (router *Router) statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := api_model.StatusResponse{
		Installed: config.SCMConfig.Installed,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		logger.Errorf("Error encoding status response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}