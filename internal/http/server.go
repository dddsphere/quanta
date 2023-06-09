package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/dddsphere/quanta/internal/infra/port/openapi"
	"github.com/dddsphere/quanta/internal/system"
)

type (
	Server struct {
		system.Worker
		http.Server
		cqrs   *CQRSHandler
		router http.Handler
	}
)

func NewServer(name string, opts ...system.Option) (server *Server) {
	return &Server{
		Worker: system.NewWorker(name, opts...),
		cqrs:   NewCQRSHadler(opts...),
	}
}

func (srv *Server) SetRouter(router http.Handler) {
	srv.router = router
}

func (srv *Server) Setup(ctx context.Context) error {
	srv.router = openapi.Handler(srv.cqrs)
	return nil
}

func (srv *Server) Start(ctx context.Context) error {
	srv.Server = http.Server{
		Addr:    srv.Address(),
		Handler: srv.router,
	}

	group, errGrpCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		srv.Log().Infof("%s started listening at %s", srv.Name(), srv.Address())
		defer srv.Log().Errorf("%s shutdown", srv.Name())

		err := srv.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	group.Go(func() error {
		<-errGrpCtx.Done()
		srv.Log().Errorf("%s shutdown", srv.Name())

		ctx, cancel := context.WithTimeout(context.Background(), srv.ShutdownTimeout())
		defer cancel()

		if err := srv.Server.Shutdown(ctx); err != nil {
			return err
		}

		return nil
	})

	return group.Wait()
}

func (srv *Server) Address() string {
	host := srv.Cfg().GetString("http.server.host")
	port := srv.Cfg().GetInt("http.server.port")
	return fmt.Sprintf("%s:%d", host, port)
}

func (srv *Server) ShutdownTimeout() time.Duration {
	secs := time.Duration(srv.Cfg().GetInt("http.server.shutdown.timeout.secs"))
	return secs * time.Second
}
