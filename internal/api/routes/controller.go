package controller

import (
	"net/http"

	"github.com/lukas-sgx/borrow/internal/api/routes/instances"
	"github.com/lukas-sgx/borrow/pkg/operatools"
)

func routesInstance(conciliator *operatools.Conciliator) http.Handler {
	endpoints := instances.NewEndpoints(conciliator)

	sub := http.NewServeMux()
	sub.HandleFunc("POST /add", endpoints.Add)
	sub.HandleFunc("DELETE /delete", endpoints.Delete)
	sub.HandleFunc("GET /{name}", endpoints.Get)

	return sub
}

func Routes(conciliator *operatools.Conciliator) http.Handler {
	basePath := "/api/v1"

	mux := http.NewServeMux()
	mux.Handle(basePath+"/instances/",
		http.StripPrefix(basePath+"/instances", routesInstance(conciliator)))

	return mux
}
