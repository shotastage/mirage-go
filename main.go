package main

import (
	"fmt"
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
			fmt.Println(context.Args().Get(0))
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
