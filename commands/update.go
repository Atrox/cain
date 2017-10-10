package commands

import (
	"github.com/atrox/cain/updater"
	"github.com/urfave/cli"
)

var UpdateCommand = &cli.Command{
	Name:   "update",
	Usage:  "update cain to the newest version (if available)",
	Action: updateAction,
}

func updateAction(c *cli.Context) error {
	err := updater.ForceRun()
	if err != nil {
		return cli.Exit(err, 1)
	}
	return nil
}
