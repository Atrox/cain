package commands

import (
	"fmt"
	"os"

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
			Name:  "path",
			Usage: "custom path",
		},
		&cli.StringFlag{
			Name:  "path-env",
			Usage: "get path from specified environment variable",
		},
		&cli.BoolFlag{
			Name:  "non-strict",
			Usage: "use non strict matching, be aware that it could match the wrong media that way",
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

	path := c.String("path")
	if c.String("path-env") != "" {
		path = os.Getenv(c.String("path-env"))
	}

	fb, err := filebot.New(conf)
	if err != nil {
		return err
	}
	fb.RetrievePath = path
	fb.NonStrictMatching = c.Bool("non-strict")

	err = fb.Execute()
	if err != nil {
		return err
	}

	fmt.Println("[+] Successfully sorted all found media files.")
	return nil
}
