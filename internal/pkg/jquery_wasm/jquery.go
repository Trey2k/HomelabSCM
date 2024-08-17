package jquery_wasm

import "syscall/js"

func Append(selector js.Value, html string) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("append", js.ValueOf(html))
}

func OnClick(selector js.Value, callback func(this js.Value, args []js.Value) interface{}) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("click", js.FuncOf(callback))
}

func Remove(selector js.Value) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("remove")
}

func SetText(selector js.Value, text string) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("text", text)
}

func GetText(selector js.Value) string {
	return js.Global().Get("jQuery").
		Invoke(selector).
		Call("text").
		String()
}

func SetValue(selector js.Value, value js.Value) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("val", value)
}

func GetValue(selector js.Value) js.Value {
	return js.Global().Get("jQuery").
		Invoke(selector).
		Call("val")
}

func Hide(selector js.Value) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("hide")
}

func Show(selector js.Value) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("show")
}

func Get(selector js.Value, toGet string) js.Value {
	return js.Global().Get("jQuery").
		Invoke(selector).
		Call(toGet)
}

func GetChildren(selector js.Value, arg string) []js.Value {
	children := js.Global().Get("jQuery").
		Invoke(selector).
		Call("children", arg)

	childrenSlice := make([]js.Value, children.Length())
	for i := 0; i < children.Length(); i++ {
		childrenSlice[i] = children.Index(i)
	}

	return childrenSlice
}

func GetParent(selector js.Value) js.Value {
	return js.Global().Get("jQuery").
		Invoke(selector).
		Call("parent")
}

func ClearChildren(selector js.Value, filter string) {
	js.Global().Get("jQuery").
		Invoke(selector).
		Call("children", filter).
		Call("remove")
}

func SetGlobalClickHandler(callback func(this js.Value, args []js.Value) any) {
	js.Global().Get("document").
		Call("addEventListener", "click", js.FuncOf(callback))
}