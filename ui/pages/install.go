package pages

import (
	"fmt"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
	"homelabscm.com/scm/internal/pkg/wasm/utils"
	"homelabscm.com/scm/ui/pages/install"
)

type Install struct {
	vgrouter.NavigatorRef

	CurrentPage vugu.Builder
}

func (c *Install) Init(ctx vugu.InitCtx) {
	fmt.Println(utils.GetCurrentPath())
	switch utils.GetCurrentPath() {
	case "/install/server_settings":
		c.CurrentPage = &install.ServerSettings{}
	default:
		c.Navigate("/install/server_settings", nil)
		c.CurrentPage = &install.ServerSettings{}
	}
}

