package cmd

import (
	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/initialization"
	"gitlab.inn4science.com/gophers/service-scaffold/workers"
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
