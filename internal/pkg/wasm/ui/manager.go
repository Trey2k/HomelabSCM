package ui

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
	"homelabscm.com/scm/internal/pkg/wasm/api"
	"homelabscm.com/scm/pkg/api_model"
	"homelabscm.com/scm/ui/pages"
)

type UIManager struct {
	apiClient *api.Client
	router    *vgrouter.Router
	status *api_model.StatusResponse
}

func NewUIManager(apiClient *api.Client) *UIManager {
	status, err := apiClient.Status()
	if err != nil {
		panic(err)
	}

	return &UIManager{
		apiClient: apiClient,
		status: status,
	}
}


// OVERALL APPLICATION WIRING IN vuguSetup
func (ui *UIManager) SetupVugu(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	// CREATE A NEW ROUTER INSTANCE
	router := vgrouter.New(eventEnv)
	ui.router = router

	// MAKE OUR WIRE FUNCTION POPULATE ANYTHING THAT WANTS A "NAVIGATOR".
	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
	})

	// CREATE THE ROOT COMPONENT
	root := &pages.Root{}
	buildEnv.WireComponent(root) // WIRE IT
	router.MustAddRouteExact("/",
		vgrouter.RouteHandlerFunc(ui.routInterceptor(func(rm *vgrouter.RouteMatch) {
			root.Body = &pages.Index{}
		})))
	// TELL THE ROUTER TO LISTEN FOR THE BROWSER CHANGING URLS
	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	// GRAB THE CURRENT BROWSER URL AND PROCESS IT AS A ROUTE
	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return root
}