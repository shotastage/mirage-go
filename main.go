package main

import (
	"fmt"
	"mirage-go/ios"
	"mirage-go/mgpkg"
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
				return nil
			}

			if context.Args().Get(0) == "mgpkg" {
				mgpkg.Bootstrap()
				return nil
			}
		} else {
			fmt.Println(context.Args().Get(0))
		}

		println("Aeguments required")

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
