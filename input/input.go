package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ValidatorFunction func(string) (interface{}, error)

func Prompt(question string, fn ValidatorFunction) interface{} {
	var out interface{}

	for {
		fmt.Printf("[?] %s: ", question)

		input, err := getInput()
		if err != nil {
			log.Fatal(err)
		}

		out, err = fn(input)
		if err == nil {
			break
		}

		fmt.Printf("[!] Error: %s\n", err.Error())
	}

	return out
}

// Create a single reader which can be called multiple times
var reader = bufio.NewReader(os.Stdin)

func getInput() (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
