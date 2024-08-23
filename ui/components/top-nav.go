package components

import (
	"github.com/vugu/vugu"
)

type TopNav struct {
	Message string
}

func (c *TopNav) Init(ctx vugu.InitCtx) {
	c.Message = "Hello, World!"
}