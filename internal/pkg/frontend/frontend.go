package frontend

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"homelabscm.com/scm/internal/pkg/config"
)

type FrontendHandler struct {
	BasePath string
	StaticFS fs.FS
	DevMode bool
}

func NewFrontendHandler(StaticFS fs.FS) *FrontendHandler {
	return &FrontendHandler{
		BasePath: config.SCMConfig.BasePath,
		StaticFS: StaticFS,
		DevMode: config.SCMConfig.DevMode,
	}
}

func (h *FrontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var staticDir http.FileSystem
	if h.DevMode {
		staticDir = http.Dir(filepath.Join(h.BasePath, "web/static"))
	} else {
		staticDir = http.FS(h.StaticFS)
	}
	log.Printf("FrontendHandler starting with dist dir: %v", staticDir)

	buildFrontend := func() (ok bool) {

		cmd := exec.Command("make", "homelab-web-client")

		cmd.Dir = h.BasePath
		b, err := cmd.CombinedOutput()
		log.Printf("make homelab-web-client - err: %v\n%s", err, b)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(500)
			fmt.Fprintf(w, "make homelab-web-client - err: %v\n%s", err, b)
			return false
		}
		return true
	}
	
	if h.DevMode && !buildFrontend() {
		return
	}

	h.serveIndex(w, r)
}

func (h *FrontendHandler) serveIndex(w http.ResponseWriter, r *http.Request) {
	in := strings.NewReader(indexHTML)
	http.ServeContent(w, r, "HomelabSCM", startupTime, in)
}

var startupTime = time.Now()

var indexHTML = `<!doctype html>
<html>
<head>
<meta charset="utf-8"/>
    <link rel="stylesheet" href="static/css/main.css">
	<link type="text/css" rel="stylesheet" href="static/css/materialize.min.css"  media="screen,projection"/>
   	<link rel="stylesheet" href="static/css/fontawesome.min.css">
<!-- styles -->
</head>
<body>
<div id="vugu_mount_point">
<img style="position: absolute; top: 50%; left: 50%;" src="https://cdnjs.cloudflare.com/ajax/libs/galleriffic/2.0.1/css/loader.gif">
</div>
<script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script> <!-- MS Edge polyfill -->
<script src="static/js/wasm_exec.js"></script>
<script type="text/javascript" src="static/js/materialize.min.js"></script>
<!-- scripts -->
<script>
var wasmSupported = (typeof WebAssembly === "object");
if (wasmSupported) {
	if (!WebAssembly.instantiateStreaming) { // polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}
	var mainWasmReq = fetch("static/wasm/homelab-web-client.wasm").then(function(res) {
		if (res.ok) {
			const go = new Go();
			WebAssembly.instantiateStreaming(res, go.importObject).then((result) => {
				go.run(result.instance);
			});		
		} else {
			res.text().then(function(txt) {
				var el = document.getElementById("vugu_mount_point");
				el.style = 'font-family: monospace; background: black; color: red; padding: 10px';
				el.innerText = txt;
			})
		}
	})
} else {
	document.getElementById("vugu_mount_point").innerHTML = 'This application requires WebAssembly support.  Please upgrade your browser.';
}
</script>
</body>
</html>`