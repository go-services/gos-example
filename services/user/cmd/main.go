package main

import (
	"github.com/go-services/gos-project/services/user/gos_gen"
	"github.com/go-services/gos-project/services/user/gos_gen/cmd"
	"os"
	"github.com/go-services/gos-project/services/user"
	"github.com/go-kit/kit/log/level"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-services/core"
)

// this is tan example of the service main function
func main() {
	// get config
	config := gos_gen.NewServiceConfig()
	// modify config here

	// create logger
	logger := makeLogger(config)

	// some initial messages
	level.Info(logger).Log("msg", fmt.Sprintf("service `%s` started", config.Name))
	defer level.Info(logger).Log("msg", fmt.Sprintf("service `%s` ended", config.Name))

	// create the service
	svc := user.NewUser(logger)

	// run the service
	cmd.Run(svc, config, logger)
}

// creates the default logger
func makeLogger(config core.ServiceConfig) log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger,
		"svc", config.Name,
		"instance", config.Instance,
		"ts", log.DefaultTimestampUTC,
		"clr", log.DefaultCaller,
	)
	return logger
}
