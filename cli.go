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

func deploy(fileMap map[string]string) error {
	for filename, sourcefile := range fileMap {
		contents := FSMustString(false, sourcefile)

		if err := generateFile(filename, contents); err != nil {
			return err
		}
	}

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
						fileMap := map[string]string{
							"main.go":   "/golang/main.go",
							"solver.go": "/golang/solver.go",
						}

						return deploy(fileMap)
					},
				},
				{
					Name:    "python",
					Aliases: []string{"p"},
					Usage:   "Initialize for a python2.7-based solution",
					Action: func(c *cli.Context) error {
						fileMap := map[string]string{
							"main.py":   "/py27/main.py",
							"solver.py": "/py27/solver.py",
						}

						return deploy(fileMap)
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
