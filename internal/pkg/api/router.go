package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	basePath string
}

func NewRouter(basePath string) *Router {
	return &Router{
		basePath: basePath,
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := mux.NewRouter()
	
	mux.HandleFunc(fmt.Sprintf("%s/status", router.basePath), 
		router.statusHandler).Methods("GET")

	mux.ServeHTTP(w, r)
}
