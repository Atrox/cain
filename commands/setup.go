package commands

import (
	"fmt"

	"github.com/atrox/cain/filebot"
	"github.com/atrox/cain/store"
	"github.com/atrox/input"
	"github.com/urfave/cli"
)

var SetupCommand = &cli.Command{
	Name:    "setup",
	Aliases: []string{"s"},
	Usage:   "create config file with sensitive defaults",
	Action:  setupAction,
}

func setupAction(c *cli.Context) error {
	_, err := filebot.Path()
	if err != nil {
		b.Println("## WARNING ##", "", "FileBot is not installed", "Cain will not work without FileBot")
	}

	conf := store.NewConfig()
	store.Get(conf)

	b.Println("Configure destinations for sorted files")

	conf.Destinations.Movie = askSaveLocation("movies", conf.Destinations.Movie)
	conf.Destinations.Series = askSaveLocation("series", conf.Destinations.Series)
	conf.Destinations.Anime = askSaveLocation("anime", conf.Destinations.Anime)

	b.Println("Configure default retrieve path for unsorted files",
		"Enter nothing to skip this step and require '--path'")

	conf.DefaultRetrievePath = askRetrievePath(conf.DefaultRetrievePath)

	b.Println("Non Strict Matching", "",
		"Should we try to match the media even if filebot is not 100% sure?",
		"Can lead to better matches and less manual work but also gets it sometimes wrong")

	conf.NonStrictMatching = input.Prompt("Enable non strict matching (y/n)", input.RequiredValidator, input.BooleanValidator).(bool)

	b.Println("Automatic Cleanup", "", "Should Cain automatically cleanup remaining unused files?")

	switch automaticCleanup := input.Prompt("Enable automatic cleanup afterwards (Y/n)", input.BooleanValidator).(type) {
	case bool:
		conf.CleanupAfterwards = automaticCleanup
	default:
		conf.CleanupAfterwards = true
	}

	b.Println("Do you want to enable automatic updates?", "",
		"If enabled and updates are available,",
		"Cain will update itself without interruptions")

	switch autoUpdate := input.Prompt("Enable automatic updates (Y/n)", input.BooleanValidator).(type) {
	case bool:
		conf.AutoUpdate = autoUpdate
	default:
		conf.AutoUpdate = true
	}

	err = store.Save(conf)
	if err != nil {
		return cli.Exit(err, 1)
	}

	b.Println("Config successfully saved",
		"You can now use 'cain run' to sort your media!",
		"", fmt.Sprintf("Location: %s", store.Path(conf)))

	return nil
}
