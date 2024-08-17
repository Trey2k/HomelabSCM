package homelab_web_client

import (
	"fmt"
	"syscall/js"

	"homelabscm.com/scm/internal/pkg/wasm_ui"
)


func StartInstaller() {
	installer_settings, err := wasm_ui.NewSettingsUI("installer-settings", js.ValueOf("#installer"))
	if err != nil {
		panic(err)
	}

	installer_settings.AddStrArrayInput("Trusted Proxies", []any{"127.0.0.1"}, "trustedProxies")
	installer_settings.AddStrArrayInput("Allowed Hosts", []any{"homelabscm.com"}, "allowedHosts")
	installer_settings.AddTextInput("Repo Storage Path", js.ValueOf("{BASE_PATH}/git-data"), "reposPath")
	installer_settings.AddTextInput("Postgres Host", js.ValueOf("127.0.0.1"), "postgresHost")
	installer_settings.AddNumberInput("Postgres Port", js.ValueOf(5432), "postgresPort")
	installer_settings.AddTextInput("Postgres User", js.ValueOf("homelab_scm"), "postgresUser")
	installer_settings.AddPasswordInput("Postgres Password", "postgresPassword")
	installer_settings.AddTextInput("Postgres Database", js.ValueOf("homelab_scm"), "postgresDatabase")
	installer_settings.AddButton("Test Connection", "testConnection", testConnection)
	installer_settings.AddButton("Install", "install", install)
}

func testConnection(this js.Value, args []js.Value) any {
	fmt.Println("Testing connection")

	return nil
}

func install(this js.Value, args []js.Value) any {
	fmt.Println("Installing")

	return nil
}
