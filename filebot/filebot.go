package filebot

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/atrox/cain/store"
)

const executableName = "filebot"

type FileBot struct {
	Path string
}

func New() (*FileBot, error) {
	path, err := exec.LookPath(executableName)
	if err != nil {
		return nil, err
	}

	return &FileBot{Path: path}, nil
}

func (f *FileBot) Execute(conf store.Config, retrievePathFlag string) error {
	// if path flag specified, overwrite retrievePath
	if retrievePathFlag != "" {
		conf.RetrievePath = retrievePathFlag
	}

	if conf.RetrievePath == "" {
		return fmt.Errorf("[!] '--path' or 'defaultRetrievePath' not specified")
	}

	commands := []string{
		"-script",
		"fn:amc",
		"--action",
		"move",
		"-non-strict",
		"--log-file",
		"amc.log",
		"--def",
		fmt.Sprintf("excludeList=%s", conf.LockFile),
		"--def",
		fmt.Sprintf("movieFormat=%s", filepath.Join(conf.Destinations.Movie, conf.NamingSchemes.Movie)),
		"--def",
		fmt.Sprintf("seriesFormat=%s", filepath.Join(conf.Destinations.Series, conf.NamingSchemes.Series)),
		"--def",
		fmt.Sprintf("animeFormat=%s", filepath.Join(conf.Destinations.Anime, conf.NamingSchemes.Anime)),
		"--def",
		fmt.Sprintf("musicFormat=%s", filepath.Join(conf.Destinations.Music, conf.NamingSchemes.Music)),
		filepath.Clean(conf.RetrievePath),
	}

	cmd := exec.Command(f.Path, commands...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
