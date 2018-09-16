package main

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-scaffold/cmd"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/info"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = info.App.Version
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()
	app.Run(os.Args)
}
