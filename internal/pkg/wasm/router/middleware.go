package router

import (
	"fmt"
	"log"
	"strings"

	"github.com/vugu/vgrouter"
)

func (r *Router) routInterceptor(handler vgrouter.RouteHandlerFunc) vgrouter.RouteHandlerFunc {
	return func(rm *vgrouter.RouteMatch) {
		if !strings.HasPrefix(rm.Path, "/install") && !r.status.Installed {
			fmt.Println("Checking status")
			status, err := r.apiClient.Status()
			if err != nil {
				log.Fatalf("failed to get status: %v", err)
				return
			}

			if !status.Installed {
				r.router.Navigate("/install", nil)
				return
			}
		}

		handler(rm)
	}
}