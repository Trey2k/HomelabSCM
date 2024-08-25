package materialize

import (
	"fmt"
	"syscall/js"
)

type Tabs struct {
	instance js.Value
}

type TabsOptions struct {
	// Duration Transition duration in milliseconds. Default: 300
	Duration int
	// OnShow Callback for when a new tab content is shown.
	OnShow func(*js.Value) any
	// Swipeable Set to true to enable swipeable tabs. This also uses the responsiveThreshold option. Default: false
	Swipeable bool
	// ResponsiveThreshold The maximum width of the screen, in pixels, where the swipeable functionality initializes. Default: Infinity
	ResponsiveThreshold int
}

func NewTabs(querySelector string, options *TabsOptions) *Tabs {
	t := &Tabs{}
	elms := js.Global().Get("document").Call("querySelectorAll", querySelector)
	tabs := js.Global().Get("M").Get("Tabs")
	fmt.Println("Tabs", tabs)
	js_options := make(map[string]any)
	if options != nil {
		if options.Duration != 0 {
			js_options["duration"] = options.Duration
		}
		if options.OnShow != nil {
			js_options["onShow"] = options.OnShow
		}
		if options.Swipeable {
			js_options["swipeable"] = options.Swipeable
		}
		if options.ResponsiveThreshold != 0 {
			js_options["responsiveThreshold"] = options.ResponsiveThreshold
		}
	}
	t.instance =  tabs.Call("init", elms, js.ValueOf(js_options))

	return t
}

func (t *Tabs) Destroy() {
	t.instance.Call("destroy")
}

func (t *Tabs) Select(tabID string) {
	t.instance.Call("select", tabID)
}