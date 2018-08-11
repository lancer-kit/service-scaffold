package main

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
)

func main() {
	cmd := cli.NewApp()
	cmd.Usage = "A " + config.ServiceName + " service"
	cmd.Version = "0.1.0"

	cmd.Commands = []cli.Command{
		serveCommand,
		migrateCommand,
	}
	cmd.Run(os.Args)
}

func initConfig(cfgPath string) {
	config.Init(cfgPath)
}

var cfgFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Value: "./config.yaml",
	},
}
