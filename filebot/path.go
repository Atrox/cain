package filebot

import "os/exec"

const executableName = "filebot"

// Path to `filebot` executable
func Path() (string, error) {
	return exec.LookPath(executableName)
}
