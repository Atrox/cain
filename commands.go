package main

import (
	"fmt"
	"strings"

	"github.com/atrox/cain/filebot"
	"github.com/atrox/cain/input"
	"github.com/atrox/cain/store"
	"github.com/urfave/cli"
)

func runCommand(c *cli.Context) error {
	conf := &store.Config{}
	err := store.Get(conf)
	if err != nil {
		return err
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

func setupCommand(c *cli.Context) error {
	fmt.Println(logo)

	if store.Exists("config") {
		fmt.Println("[!] Configuration File already exists. Config will get overwritten!")
	}

	fmt.Println("[+] You can skip steps if you just insert nothing at all")

	conf := store.NewConfig()
	conf.Destinations.Movie = saveToPrompt("movies")
	conf.Destinations.Series = saveToPrompt("series")
	conf.Destinations.Anime = saveToPrompt("anime")
	conf.RetrievePath = retrievePath()

	err := store.Save(conf)
	if err != nil {
		return err
	}

	fmt.Println("[+] Config successfully saved")
	fmt.Println("[+] You can now use `run` to sort your media!")
	return nil
}

func updateCommand(c *cli.Context) error {
	return equinoxUpdate()
}

func saveToPrompt(name string) string {
	p := input.Prompt(fmt.Sprintf("Where to put the sorted %s", name), input.PathValidator).(string)

	if p == "" {
		fmt.Printf("[+] %s will get ignored by Cain\n", strings.Title(name))
	} else {
		fmt.Printf("[+] %s will get saved to %s\n", strings.Title(name), p)
	}

	return p
}

func retrievePath() string {
	p := input.Prompt("Where should I look for media if '--path' is not specified", input.PathValidator).(string)

	if p == "" {
		fmt.Println("[+] Default Value for '--path' not specified. Parameter is now required.")
	} else {
		fmt.Printf("[+] Default Value for '--path' set as %s.\n", p)
	}

	return p
}
