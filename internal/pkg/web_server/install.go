package web_server

import (
	"net/http"

	"homelabscm.com/scm/internal/pkg/config"
	"homelabscm.com/scm/internal/pkg/logger"
)

func installHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":  "HomeLab SCM - Installer",
		"Config": config.SCMConfig,
	}

	err := executeTemplate("install", w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Errorf(err.Error())
		return
	}
}
