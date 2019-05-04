package main

import (
	"os"

	"github.com/lancer-kit/service-scaffold/cmd"
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/lancer-kit/service-scaffold/info"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = info.App.Version
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()
	app.Run(os.Args)
}
