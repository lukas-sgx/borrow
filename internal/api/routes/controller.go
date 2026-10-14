package controller

import (
	"net/http"

	"github.com/lukas-sgx/borrow/internal/api/routes/instances"
	"github.com/lukas-sgx/borrow/pkg/operatools"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func routesInstance(conciliator *operatools.Conciliator, client client.Client) http.Handler {
	endpoints := instances.NewEndpoints(conciliator, client)

	sub := http.NewServeMux()
	sub.HandleFunc("POST /add", endpoints.Add)
	sub.HandleFunc("DELETE /delete", endpoints.Delete)
	sub.HandleFunc("GET /{name}", endpoints.Get)

	return sub
}

func Routes(conciliator *operatools.Conciliator, client client.Client) http.Handler {
	basePath := "/api/v1"

	mux := http.NewServeMux()
	mux.Handle(basePath+"/instances/",
		http.StripPrefix(basePath+"/instances", routesInstance(conciliator, client)))

	return mux
}
