package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize the current folder for a new CJam problem set",
			Subcommands: []cli.Command{
				{
					Name:    "golang",
					Aliases: []string{"go"},
					Usage:   "Initialize for a golang-based solution",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:    "python",
					Aliases: []string{"p"},
					Usage:   "Initialize for a python2.7-based solution",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
