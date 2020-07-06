package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"lancer-kit/service-scaffold/cmd"
	"lancer-kit/service-scaffold/config"
)

func main() {
	fmt.Printf("%+v \n", config.AppInfo())

	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = config.AppInfo().Version
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
