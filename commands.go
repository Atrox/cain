package main

import (
	"fmt"

	"github.com/atrox/box"
	"github.com/atrox/cain/filebot"
	"github.com/atrox/cain/input"
	"github.com/atrox/cain/store"
	"github.com/urfave/cli"
)

var b = box.New()

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
	conf := store.NewConfig()
	store.GetOrCreate(conf)

	b.Println("Configure destinations for sorted files")

	conf.Destinations.Movie = askSaveLocation("movies", conf.Destinations.Movie)
	conf.Destinations.Series = askSaveLocation("series", conf.Destinations.Series)
	conf.Destinations.Anime = askSaveLocation("anime", conf.Destinations.Anime)

	b.Println("Configure default retrieve path for unsorted files",
		"Enter nothing to skip this step and require '--path'")

	conf.RetrievePath = retrievePath(conf.RetrievePath)

	b.Println("Do you want to enable automatic updates?", "",
		"If enabled and updates are available,",
		"Cain will update itself without interruptions")

	conf.AutoUpdate = input.Prompt("Enable automatic updates (Y/n)", input.BooleanValidator(true)).(bool)

	err := store.Save(conf)
	if err != nil {
		return err
	}

	b.Println("Config successfully saved",
		"You can now use 'cain run' to sort your media!")

	return nil
}

func updateCommand(c *cli.Context) error {
	return appUpdater.ForceRun()
}

func askSaveLocation(name string, current string) string {
	if current != "" {
		prompt := input.Prompt(fmt.Sprintf("Where to put the sorted %s (default=%s)", name, current), input.PathValidator(true)).(string)
		if prompt == "" {
			return current
		}

		return prompt
	}

	prompt := input.Prompt(fmt.Sprintf("Where to put the sorted %s", name), input.PathValidator(false)).(string)
	return prompt
}

func retrievePath(current string) string {
	var text string
	if current == "" {
		text = "Default retrieve path"
	} else {
		text = fmt.Sprintf("Default retrieve path (default=%s)", current)
	}

	prompt := input.Prompt(text, input.PathValidator(true)).(string)

	if prompt == "" && current != "" {
		fmt.Printf("[+] Default Value for '--path' set as %s.\n", current)
		return current
	} else if prompt == "" {
		fmt.Println("[+] Default Value for '--path' not specified. Parameter is now required.")
	} else {
		fmt.Printf("[+] Default Value for '--path' set as %s.\n", prompt)
	}

	return prompt
}
