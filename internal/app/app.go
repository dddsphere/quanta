package app

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"google.golang.org/grpc"

	//h "github.com/dddsphere/quanta/internal/driver/http"
	"github.com/dddsphere/quanta/internal/config"
	"github.com/dddsphere/quanta/internal/core/errors"
	"github.com/dddsphere/quanta/internal/log"
	"github.com/dddsphere/quanta/internal/system"
)

type App struct {
	sync.Mutex
	system.Worker
	opts       []system.Option
	supervisor system.Supervisor
	http       *http.Server
}

func NewApp(name, namespace string, log log.Logger) (app *App) {
	cfg := config.Load(namespace)

	opts := []system.Option{
		system.WithConfig(cfg),
		system.WithLogger(log),
	}

	app = &App{
		Worker: system.NewWorker(name, opts...),
		opts:   opts,
	}

	return app
}

func (app *App) Run() (err error) {
	ctx := context.Background()

	err = app.Setup(ctx)
	if err != nil {
		return errors.Wrap(runError, err)
	}

	return app.Start(ctx)
}

func (app *App) Setup(ctx context.Context) error {
	app.EnableSupervisor()

	// Setup db connections

	// Setup repos

	// Setup services

	// Setup http handlers

	// Setup gRPC servers

	// Setup event bus

	return nil
}

func (app *App) Start(ctx context.Context) error {
	app.Log().Infof("%s starting...", app.Name())
	defer app.Log().Infof("%s stopped", app.Name())

	app.supervisor.AddTasks(
	//app.http.Start,
	//app.grpc.Start,
	)

	app.Log().Infof("%s started!", app.Name())

	return app.supervisor.Wait()
}

func (app *App) Stop(ctx context.Context) error {
	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	return nil
}

func (app *App) EnableSupervisor() {
	name := fmt.Sprintf("%s-supervisor", app.Name())
	app.supervisor = system.NewSupervisor(name, true, app.opts)
}

// Service interface

func (app *App) RegisterHTTPHandler(http.Handler) {
	app.Log().Infof("No registered HTTP handlers for %s", app.Name())
}

func (app *App) RegisterGRPCServer(srv *grpc.Server) {
	app.Log().Infof("No registered gRPC servers for %s", app.Name())
}
