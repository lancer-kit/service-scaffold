package cmd

import (
	"lancer-kit/service-scaffold/cmd/modules"
	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/workers"

	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli"
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
	cfg, err := modules.Init(c)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	logger := log.Get().WithField("app", config.ServiceName)
	workers.InitChief(logger, cfg).Run()

	return nil
}
