package pages

import (
	"fmt"

	"github.com/vugu/vugu"
)

type Index struct {
	Message string
}

func (c *Index) Init(ctx vugu.InitCtx) {
	c.Message = "Hello, World!"
}

func (c *Index) Submit(event vugu.DOMEvent) {
	fmt.Println("Submit clicked")
	event.PreventDefault()
}