package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/atrox/cain/commands"
	"github.com/atrox/cain/store"
	"github.com/atrox/cain/updater"
	"github.com/urfave/cli"
)

var (
	version    = "master"
	appUpdater *updater.Updater
)

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

var logo = strings.TrimLeft(fmt.Sprintf(`
          _             _                    _          _          
        /\ \           / /\                 /\ \       /\ \     _  
       /  \ \         / /  \                \ \ \     /  \ \   /\_\
      / /\ \ \       / / /\ \               /\ \_\   / /\ \ \_/ / /
     / / /\ \ \     / / /\ \ \             / /\/_/  / / /\ \___/ / 
    / / /  \ \_\   / / /  \ \ \           / / /    / / /  \/____/  
   / / /    \/_/  / / /___/ /\ \         / / /    / / /    / / /   
  / / /          / / /_____/ /\ \       / / /    / / /    / / /    
 / / /________  / /_________/\ \ \  ___/ / /__  / / /    / / /     
/ / /_________\/ / /_       __\ \_\/\__\/_/___\/ / /    / / /      
\/____________/\_\___\     /____/_/\/_________/\/_/     \/_/       
============================================================
 Version %s                                 by Atrox.ME
`, version), "\n")
