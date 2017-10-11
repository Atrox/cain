package commands

import (
	"fmt"

	"github.com/atrox/cain/updater"
	"github.com/urfave/cli"
)

var UpdateCommand = &cli.Command{
	Name:   "update",
	Usage:  "update cain to the newest version (if available)",
	Action: updateAction,
}

func updateAction(c *cli.Context) error {
	appUpdater := updater.New(true)
	updated, err := appUpdater.Run()
	if err != nil {
		return cli.Exit(err, 1)
	}
	if !updated {
		b.Println("You are on the newest version", "", fmt.Sprintf("Version %s", updater.Version))
	}
	return nil
}
