package controller

import (
	"net/http"

	"github.com/lukas-sgx/borrow/internal/api/routes/instances"
)

func routesInstance() http.Handler {
	sub := http.NewServeMux()
	sub.HandleFunc("POST /add", instances.Add)
	sub.HandleFunc("DELETE /delete", instances.Delete)
	sub.HandleFunc("GET /{name}", instances.Get)

	return sub
}

func Routes() http.Handler {
	basePath := "/api/v1"

	mux := http.NewServeMux()
	mux.Handle(basePath+"/instances/",
		http.StripPrefix(basePath+"/instances", routesInstance()))

	return mux
}
