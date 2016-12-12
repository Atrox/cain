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
	RetrievePath string

	exePath string
	config  *store.Config
	cmds    []string
}

func New(conf *store.Config) (*FileBot, error) {
	path, err := exec.LookPath(executableName)
	if err != nil {
		return nil, err
	}

	// default commands
	commands := []string{
		"-script", "fn:amc",
		"--action", "move",
		"--log-file", "amc.log",
		"--def", "clean=y",
		"-non-strict",
	}

	return &FileBot{
		exePath: path,
		config:  conf,
		cmds:    commands,
	}, nil
}

func (f *FileBot) Execute() error {
	retrievePath := f.config.RetrievePath

	// if path flag specified, overwrite retrievePath
	if f.RetrievePath != "" {
		retrievePath = f.RetrievePath
	}

	if retrievePath == "" {
		return fmt.Errorf("[!] '--path' or 'defaultRetrievePath' not specified")
	}

	f.defCmds("excludeList", "movieFormat", "seriesFormat", "animeFormat", "musicFormat")
	f.cmds = append(f.cmds, filepath.Clean(retrievePath))

	cmd := exec.Command(f.exePath, f.cmds...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (f *FileBot) defCmds(names ...string) {
	for _, name := range names {
		f.defCmd(name)
	}
}

func (f *FileBot) defCmd(name string) {
	loc := f.getPath(name)
	if loc == "" {
		return
	}

	f.cmds = append(f.cmds, "--def", fmt.Sprintf("%s=%s", name, loc))
}

func (f *FileBot) getPath(name string) string {
	switch name {
	case "excludeList":
		if f.config.LockFile == "" {
			return ""
		}

		return filepath.Clean(f.config.LockFile)
	case "movieFormat":
		if f.config.Destinations.Movie == "" {
			return ""
		}

		return filepath.Join(f.config.Destinations.Movie, f.config.NamingSchemes.Movie)
	case "seriesFormat":
		if f.config.Destinations.Series == "" {
			return ""
		}

		return filepath.Join(f.config.Destinations.Series, f.config.NamingSchemes.Series)
	case "animeFormat":
		if f.config.Destinations.Anime == "" {
			return ""
		}

		return filepath.Join(f.config.Destinations.Anime, f.config.NamingSchemes.Anime)
	default:
		return ""
	}
}
