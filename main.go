package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/inn4sc/go-skeleton/config"
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

func initModules(cfgPath string) {
	config.Init(cfgPath)
	config.InitLog()
	config.InitDB()
}

var cfgFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Value: "./config.yaml",
	},
}
