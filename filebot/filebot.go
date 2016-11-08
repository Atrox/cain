package filebot

import (
	"fmt"
	"os/exec"

	"github.com/atrox/cain/store"

	"os"
)

const Executable = "filebot"

type FileBot struct {
	Path string
}

func New() (*FileBot, error) {
	path, err := exec.LookPath(Executable)
	if err != nil {
		return nil, err
	}

	return &FileBot{Path: path}, nil
}

func (f *FileBot) Execute(conf *store.Config, retrievePathFlag string) error {
	// set retrievePath to defaultRetrievePath from config
	retrievePath := conf.RetrievePath

	// if path flag specified, overwrite retrievePath
	if retrievePathFlag != "" {
		retrievePath = retrievePathFlag
	}

	if retrievePath == "" {
		return fmt.Errorf("[!] '--path' or 'defaultRetrievePath' not specified")
	}

	cmd := exec.Command(Executable, "-help")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
