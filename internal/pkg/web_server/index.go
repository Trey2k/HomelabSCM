package web_server

import (
	"net/http"

	"homelabscm.com/scm/internal/pkg/logger"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "HomeLab SCM",
	}

	err := executeTemplate("index", w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Errorf(err.Error())
		return
	}
}
