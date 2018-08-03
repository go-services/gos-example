package cmd

import (
	"github.com/go-services/gos-project/services/user/gos_gen"
	"github.com/gorilla/mux"
	gosHttp "github.com/go-services/gos-project/services/user/gos_gen/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"fmt"
	"github.com/oklog/run"
	"net/http"
	"net"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-services/gos-project/services/user/config"
)

// Run the service
func Run(svc gos_gen.UserService, cnf *config.ServiceConfig, logger log.Logger) {
	var g run.Group
	{
		// setup HTTP transport
		listener, router := setupHttp(svc, cnf, logger)
		g.Add(func() error {
			return http.Serve(listener, router)
		}, func(error) {
			listener.Close()
		})
	}
	{
		// TODO add GRPC HERE
	}
	{
		// handle cancel interrupt
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	// run the group
	level.Error(logger).Log("exit", g.Run())
}

// sets up the http transport for all the resources that have methods that use http
func setupHttp(service gos_gen.UserService, cnf *config.ServiceConfig, logger log.Logger) (net.Listener, *mux.Router) {
	router := mux.NewRouter()
	listener, err := net.Listen("tcp", cnf.HTTPAddress)

	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// setup base resource http
	setupBaseResourceHttp(service.Base().Http(), router)

	return listener, router
}

func setupBaseResourceHttp(b gosHttp.BaseHttp, router *mux.Router) {
	for _, ep := range b.Iterate() {
		for _, route := range ep.MethodRoutes() {
			router.
				Methods(string(route.Method)).
				Path(route.Route).
				Handler(b.Index().Handler()).
				Name(route.Name)
		}
	}
}
