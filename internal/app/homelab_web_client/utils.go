package homelab_web_client

import "syscall/js"

func getURLPath() string {
	location := js.Global().Get("window").Get("location")
	pathname := location.Get("pathname").String()
	return pathname
}
