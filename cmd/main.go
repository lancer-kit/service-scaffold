package cmd

import (
	"lancer-kit/service-scaffold/config"

	"github.com/lancer-kit/uwe/v2"
	"github.com/urfave/cli"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		migrateCmd(),
		serveCmd(),
		uwe.CliCheckCommand(config.AppInfo(), func(c *cli.Context) []uwe.WorkerName {
			cfg, _ := config.ReadConfig(c.GlobalString(config.FlagConfig))
			cfg.FillDefaultWorkers()
			res := make([]uwe.WorkerName, len(cfg.Workers))
			for i, worker := range cfg.Workers {
				res[i] = uwe.WorkerName(worker)
			}
			return res
		}),
	}
}

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  config.FlagConfig + ", c",
			Value: "./config.yaml",
		},
	}
}
