package cmd

import "github.com/urfave/cli"

func GetAll() []cli.Command {
	return []cli.Command{
		migrateCommand,
		serveCommand,
	}
}

var cfgFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Value: "./config.yaml",
	},
}
