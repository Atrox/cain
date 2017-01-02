package commands

import (
	"fmt"

	"github.com/atrox/cain/filebot"
	"github.com/atrox/cain/store"
	"github.com/urfave/cli"
)

var RunCommand = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "run cain",
	Action:  runAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "path, p",
			Usage: "custom path",
		},
	},
}

func runAction(c *cli.Context) error {
	conf := &store.Config{}
	err := store.Get(conf)
	if err != nil {
		b.Println("Configuration not found", "Starting 'cain setup'")
		return setupAction(c)
	}

	fb, err := filebot.New(conf)
	if err != nil {
		return err
	}
	fb.RetrievePath = c.String("path")

	err = fb.Execute()
	if err != nil {
		return err
	}

	fmt.Println("[+] Successfully sorted all found media files.")
	return nil
}
