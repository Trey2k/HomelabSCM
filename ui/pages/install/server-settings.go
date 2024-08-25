package install

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type ServerSettings struct {
	vgrouter.NavigatorRef
}

func (c *ServerSettings) Init(ctx vugu.InitCtx) {
}