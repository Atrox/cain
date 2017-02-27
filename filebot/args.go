package filebot

import "fmt"

type args []string

func newArgs() *args {
	return &args{
		"-script", "fn:amc",
		"--action", "move",
	}
}

func (a *args) Add(s ...string) {
	*a = append(*a, s...)
}

func (a *args) AddDefinition(name, value string) {
	if name == "" || value == "" {
		return
	}

	a.Add("--def", fmt.Sprintf("%s=%s", name, value))
}
