package homelab_scm

import (
	"homelabscm.com/scm/internal/pkg/web_server"
)

func Run() error {
	return web_server.Run()
}
