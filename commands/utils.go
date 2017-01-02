package commands

import (
	"fmt"

	"github.com/atrox/box"
	"github.com/atrox/cain/input"
)

var b = box.New()

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

func askRetrievePath(current string) string {
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
