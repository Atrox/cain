package filebot

import "os/exec"

const executableName = "filebot"

// Path looks for `filebot` executable
func Path() (string, error) {
	path, err := exec.LookPath(executableName)
	return path, err
}
