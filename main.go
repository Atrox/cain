package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:      "cain",
		HelpName:  "cain",
		Usage:     "simple media management",
		Version:   "v0.1.2",
		Authors:   []*cli.Author{{Name: "Atrox", Email: "mail@atrox.me"}},
		Copyright: "(c) 2016 Atrox (Atrox.ME)",
		Commands: []*cli.Command{
			{
				Name:    "setup",
				Aliases: []string{"s"},
				Usage:   "setup cain...",
				Action:  setupCommand,
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run cain on this",
				Action:  runCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "path, p",
						Usage: "path to use for run",
					},
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update cain to the newest version",
				Action:  updateCommand,
			},
		},
		Before: before,
		After:  after,
	}

	app.Run(os.Args)
}

func before(c *cli.Context) error {
	checkForUpdates()

	return nil
}

func after(c *cli.Context) error {
	msg := <-updatesChan
	if msg != "" {
		fmt.Println(msg)
	}

	return nil
}
