$(function() {
    const go = new Go();
        WebAssembly.instantiateStreaming(fetch("/static/wasm/homelab-web-client.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
});