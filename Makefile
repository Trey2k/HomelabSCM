all: build-scm build-scm-shell debug

check-version-file:
	./scripts/check-version.sh

homelab-web-client:
	go generate ./ui
	GOOS=js GOARCH=wasm go build -o web/static/wasm/homelab-web-client.wasm wasm/homelab-web-client/main.go

build-scm: homelab-web-client
	GOOS=linux GOARCH=amd64 go build -o build/bin/homelab-scm cmd/homelab-scm/main.go

build-scm-shell:
	GOOS=linux GOARCH=amd64 go build -o build/bin/homelab-scm-shell cmd/homelab-scm-shell/main.go

package-debian: check-version-file
	./scripts/package-debian.sh

debug:
	./scripts/debug.sh

clean:
	go clean
	rm build/bin/homelab-scm
	rm build/bin/homelab-scm-shell
	rm build/bin/homelab-scm.deb
	rm internal/web/static/wasm/homelab-web-client.wasm
	rm -rf build/tmp

.PHONY: all build-scm build-scm-shell package-debian clean
