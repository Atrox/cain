package main

import (
	"fmt"
	"os"

	"github.com/atrox/cain/commands"
	"github.com/atrox/cain/store"
	"github.com/atrox/cain/updater"
	"github.com/urfave/cli"
)

var version = "master"

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
		Version:  version,
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

	// check only for updates if binary is versioned
	if version != "master" {
		// get configuration
		conf := new(store.Config)
		store.Get(conf)

		// check for updates in background
		appUpdater = updater.New(conf.AutoUpdate)
	}

	return nil
}

func after(c *cli.Context) error {
	if appUpdater == nil {
		return nil
	}

	return appUpdater.Run()
}
