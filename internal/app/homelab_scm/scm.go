package homelab_scm

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"homelabscm.com/scm/internal/pkg/config"
	"homelabscm.com/scm/internal/pkg/frontend"
	"homelabscm.com/scm/internal/pkg/logger"
	web_fs "homelabscm.com/scm/web"
)

func Run() error {
	r := mux.NewRouter()
	http.Handle("/", httpInterceptor(r))

	frontend := frontend.NewFrontendHandler(web_fs.StaticFS)

//	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(web_fs.StaticFS)))

	r.PathPrefix("/").Handler(frontend)

	conn := fmt.Sprintf(":%d", config.SCMConfig.Port)
	err := http.ListenAndServe(conn, nil)
	if err != nil {
		return err
	}

	logger.Infof("Server started on port %d", config.SCMConfig.Port)

	return nil
}
