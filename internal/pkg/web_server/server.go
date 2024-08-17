package web_server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"homelabscm.com/scm/internal/pkg/config"
	web_fs "homelabscm.com/scm/web"
)

func Run() error {
	r := mux.NewRouter()
	http.Handle("/", httpInterceptor(r))

	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(web_fs.StaticFS)))

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/install", installHandler)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.SCMConfig.Port), nil)
}
