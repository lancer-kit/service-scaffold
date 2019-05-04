package cmd

import (
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/lancer-kit/service-scaffold/initialization"
	"github.com/lancer-kit/service-scaffold/workers"
	"github.com/urfave/cli"
)

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	cfg := initialization.Init(c)

	workers.GetChief().RunAll(cfg.Log.AppName, cfg.Workers...)
	return nil
}
