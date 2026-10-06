package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Server struct {
	Addr   string
	Client client.Client
}

func (s *Server) NeedLeaderElection() bool {
	return false
}

func (s *Server) shutdownServer(ctx context.Context, srv *http.Server) {
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}

func (s *Server) routesNamespace() http.Handler {
	sub := http.NewServeMux()
	sub.HandleFunc("POST /add", s.CreateNamespace)

	return sub
}

func (s *Server) routes() http.Handler {
	basePath := "/api/v1"

	mux := http.NewServeMux()
	mux.Handle(basePath+"/namespaces/",
		http.StripPrefix(basePath+"/namespaces", s.routesNamespace()))

	return mux
}

func (s *Server) Start(ctx context.Context) error {

	srv := &http.Server{Addr: s.Addr, Handler: s.routes()}

	go s.shutdownServer(ctx, srv)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) CreateNamespace(response http.ResponseWriter, req *http.Request) {
	response.Header().Add("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
}

func NewServer(addr string, c client.Client) *Server {
	return &Server{
		Addr:   addr,
		Client: c,
	}
}
