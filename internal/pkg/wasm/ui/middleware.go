package ui

import (
	"fmt"
	"log"

	"github.com/vugu/vgrouter"
)

func (ui *UIManager) routInterceptor(handler vgrouter.RouteHandlerFunc) vgrouter.RouteHandlerFunc {
	return func(rm *vgrouter.RouteMatch) {
		if rm.Path != "/install" && !ui.status.Installed {
			fmt.Println("Checking status")
			status, err := ui.apiClient.Status()
			if err != nil {
				log.Fatalf("failed to get status: %v", err)
				return
			}

			if !status.Installed {
				ui.router.Navigate("/install", nil)
				return
			}
		}

		handler(rm)
	}
}