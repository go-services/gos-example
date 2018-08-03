package main

import (
	"github.com/go-services/gos-project/services/user/gos_gen/cmd"
	"os"
	"github.com/go-services/gos-project/services/user"
	"github.com/go-kit/kit/log/level"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-services/gos-project/services/user/db"
	"github.com/go-services/gos-project/services/user/config"
)

// this is tan example of the service main function
func main() {
	// get config
	cnf := config.Get()
	// modify config here

	// create logger
	logger := makeLogger(cnf)

	// some initial messages
	level.Info(logger).Log("msg", fmt.Sprintf("service `%s` started", cnf.Name))
	defer level.Info(logger).Log("msg", fmt.Sprintf("service `%s` ended", cnf.Name))

	// create the service
	svc := user.NewUser(logger)

	db.Session()
	defer db.Close()
	// run the service
	cmd.Run(svc, cnf, logger)
}

// creates the default logger
func makeLogger(config *config.ServiceConfig) log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger,
		"svc", config.Name,
		"instance", config.Instance,
		"ts", log.DefaultTimestampUTC,
		"clr", log.DefaultCaller,
	)
	return logger
}
