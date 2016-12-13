package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const Version = "v0.1.3"

var app *cli.App

func main() {
	app = &cli.App{
		Name:     "cain",
		HelpName: "cain",
		Usage:    "automated media management",
		Version:  Version,
		Authors:  []*cli.Author{{Name: "Atrox", Email: "mail@atrox.me"}},
		Commands: []*cli.Command{
			{
				Name:    "setup",
				Aliases: []string{"s"},
				Usage:   "create config file with sensitive defaults",
				Action:  setupCommand,
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run cain",
				Action:  runCommand,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "path, p",
						Usage: "custom path",
					},
				},
			},
			{
				Name:   "update",
				Usage:  "update cain to the newest version (if available)",
				Action: updateCommand,
			},
		},
		Before: before,
		After:  after,
	}

	app.Run(os.Args)
}

func before(c *cli.Context) error {
	// print logo
	fmt.Println(logo)

	// check for updates in background
	checkForUpdates()

	return nil
}

func after(c *cli.Context) error {
	// wait for update request to finish
	// and print message (if necessary)
	msg := <-updatesChan
	if msg != "" {
		fmt.Println(msg)
	}

	return nil
}
