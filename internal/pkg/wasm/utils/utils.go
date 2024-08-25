package utils

import "syscall/js"

var (
	window = js.Global().Get("window")
	document = window.Get("document")
)

func JSGet(key string) js.Value {
	return js.Global().Get(key)
}

func GetCurrentPath() string {
	return window.Get("location").Get("pathname").String()
}

func QuerySelectorAll(selector string) js.Value {
	return document.Call("querySelectorAll", selector)
}

func QuerySelector(selector string) js.Value {
	return document.Call("querySelector", selector)
}

func AddEventListener(event string, callback func()) {
	window.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback()
		return nil
	}))
}