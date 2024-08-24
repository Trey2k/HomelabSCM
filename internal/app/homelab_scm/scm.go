package homelab_scm

import (
	"fmt"
	"mime"
	"net/http"
	"path"

	"github.com/gorilla/mux"

	"homelabscm.com/scm/internal/pkg/api"
	"homelabscm.com/scm/internal/pkg/config"
	"homelabscm.com/scm/internal/pkg/frontend"
	"homelabscm.com/scm/internal/pkg/logger"
	web_fs "homelabscm.com/scm/web"
)

func Run() error {
	r := mux.NewRouter()
	http.Handle("/", httpInterceptor(r))

	frontend := frontend.NewFrontendHandler(web_fs.StaticFS)

	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".wasm", "application/wasm")

	if config.SCMConfig.DevMode {
		logger.Infof("Running in development mode")
		r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(path.Join(config.SCMConfig.BasePath, "web"))))
	} else {
		r.PathPrefix("/static/").Handler(http.FileServer(http.FS(web_fs.StaticFS)))
	}

	api_router := api.NewRouter("/api/v1")
	r.PathPrefix("/api/v1").Handler(api_router)

	r.PathPrefix("/").Handler(frontend)

	conn := fmt.Sprintf(":%d", config.SCMConfig.Port)
	err := http.ListenAndServe(conn, nil)
	if err != nil {
		return err
	}

	logger.Infof("Server started on port %d", config.SCMConfig.Port)

	return nil
}
