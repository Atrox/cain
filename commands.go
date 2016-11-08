package main

import (
	"fmt"

	"github.com/atrox/cain/filebot"
	"github.com/atrox/cain/input"
	"github.com/atrox/cain/store"
	"github.com/urfave/cli"
)

func runCommand(c *cli.Context) error {
	fb, err := filebot.New()
	if err != nil {
		return err
	}

	conf := &store.Config{}

	err = store.Get(conf)
	if err != nil {
		return err
	}

	err = fb.Execute(conf, c.String("path"))
	if err != nil {
		return err
	}

	fmt.Println(conf)
	fmt.Println(fb)

	return nil
}

func setupCommand(c *cli.Context) error {
	fmt.Println(logo)

	if store.Exists("config") {
		fmt.Println("[!] Configuration File already exists. Config will get overwritten!")
	}

	conf := store.NewConfig()
	conf.Destinations.Movie = saveToPrompt("movies")
	conf.Destinations.Series = saveToPrompt("series")
	conf.Destinations.Anime = saveToPrompt("anime")
	conf.Destinations.Music = saveToPrompt("music")
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
	p := input.Prompt(fmt.Sprintf("Where to put the sorted %s", name), input.PathValidator(false))
	fmt.Printf("[+] %s will get saved to %s\n", name, p)

	return p.(string)
}

func retrievePath() string {
	p := input.Prompt("Where should I look for media if '--path' is not specified (enter nothing to skip)", input.PathValidator(true)).(string)

	if p == "" {
		fmt.Println("[+] Default Value for '--path' not specified. Parameter is now required.")
	} else {
		fmt.Printf("[+] Default Value for '--path' set as %s.\n", p)
	}

	return p
}
