package main

import (
	"os"

	"github.com/urfave/cli"
)

func generateFile(filename string, contents string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(contents)

	return nil
}

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
						templateFiles := map[string]string{
							"main.go":   "/golang/main.go",
							"solver.go": "/golang/solver.go",
						}

						for filename, sourcefile := range templateFiles {
							contents := FSMustString(false, sourcefile)

							if err := generateFile(filename, contents); err != nil {
								return err
							}
						}

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
