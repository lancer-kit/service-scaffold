package cmd

import (
	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/initialization"
	"lancer-kit/service-scaffold/workers"
)

func serveCmd() cli.Command {
	var serveCommand = cli.Command{
		Name:   "serve",
		Usage:  "starts " + config.ServiceName + " workers",
		Action: serveAction,
	}
	return serveCommand
}

func serveAction(c *cli.Context) error {
	cfg := initialization.Init(c)

	logger := log.Get().WithField("app", config.ServiceName)

	chief := workers.InitChief(logger, cfg)
	chief.Run()
	return nil
}
