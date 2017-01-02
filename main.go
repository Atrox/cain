package main

import (
	"fmt"
	"os"

	"github.com/atrox/cain/commands"
	"github.com/atrox/cain/store"
	"github.com/atrox/cain/updater"
	"github.com/urfave/cli"
)

const Version = "v0.2.0"

func main() {
	cmds := []*cli.Command{
		commands.SetupCommand,
		commands.RunCommand,
		commands.UpdateCommand,
	}

	app := &cli.App{
		Name:     "cain",
		HelpName: "cain",
		Usage:    "automated media management",
		Version:  Version,
		Authors:  []*cli.Author{{Name: "Atrox", Email: "mail@atrox.me"}},
		Commands: cmds,
		Before:   before,
		After:    after,
	}

	app.Run(os.Args)
}

var appUpdater *updater.Updater

func before(c *cli.Context) error {
	fmt.Println(logo)

	// get configuration
	conf := &store.Config{}
	store.Get(conf)

	// check for updates in background
	appUpdater = updater.New(conf.AutoUpdate)

	return nil
}

func after(c *cli.Context) error {
	// wait for update request to finish
	err := appUpdater.Run()
	if err != nil {
		return err
	}

	return nil
}
