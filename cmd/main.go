package cmd

import "github.com/urfave/cli"

func GetCommands() []cli.Command {
	return []cli.Command{
		migrateCommand,
		serveCommand,
	}
}

const FlagConfig = "config"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  FlagConfig + ", c",
			Value: "./config.yaml",
		},
	}
}
