package main

import (
	"fmt"
	"mirage/ios"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "MIRAGE"
	app.Usage = "Dev tool"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("bootstrap") {
			if context.Args().Get(0) == "ios" {
				ios.Bootstrap()
			}
		} else {
			fmt.Println(context.Args().Get(0))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "bootstrap, bs",
			Usage: "Bootstrap mirage project",
		},
	}

	app.Run(os.Args)
}
