package main

import (
	"fmt"
	"mirage-go/gormdb"
	"mirage-go/ios"
	"mirage-go/mgpkg"
	"mirage-go/touch"

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

			if context.Args().Get(0) == "gorm" {
				gormdb.Bootstrap()
				return nil
			}
		} else {
			fmt.Println(context.Args().Get(0))
		}

		println("Aeguments required")

		return nil
	}

	app.Action = func(context *cli.Context) error {

		touch.Touch("py")

		println("Aeguments required")

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "bootstrap, bs",
			Usage: "Bootstrap mirage project",
		},
		cli.BoolFlag{
			Name:  "touch, cr",
			Usage: "Create new text file with copyright header",
		},
	}

	app.Run(os.Args)
}
