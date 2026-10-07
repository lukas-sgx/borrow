package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	controller "github.com/lukas-sgx/borrow/internal/api/routes"
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

func (s *Server) Start(ctx context.Context) error {

	srv := &http.Server{Addr: s.Addr, Handler: controller.Routes()}

	go s.shutdownServer(ctx, srv)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func NewServer(addr string, c client.Client) *Server {
	return &Server{
		Addr:   addr,
		Client: c,
	}
}
