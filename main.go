package main

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-scaffold/cmd"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = "0.1.0"

	app.Commands = cmd.GetAll()
	app.Run(os.Args)
}
