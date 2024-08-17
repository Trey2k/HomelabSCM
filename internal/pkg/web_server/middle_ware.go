package web_server

import (
	"net/http"
	"strings"

	"homelabscm.com/scm/internal/pkg/config"
)

func httpInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !config.SCMConfig.Installed && r.URL.Path != "/install" && !strings.Contains(r.URL.Path, "/static/") {
			http.Redirect(w, r, "/install", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
