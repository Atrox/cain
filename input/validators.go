package input

import (
	"fmt"
	"os"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
)

func PathValidator(allowEmpty bool) func(input string) (interface{}, error) {
	return func(input string) (interface{}, error) {
		if input == "" {
			if allowEmpty {
				return input, nil
			}
			return nil, fmt.Errorf("This option is required!")
		}

		dir, err := homedir.Expand(input)
		if err != nil {
			return nil, err
		}

		info, err := os.Stat(dir)
		if err != nil {
			return nil, err
		}

		if !info.IsDir() {
			return nil, fmt.Errorf("%s is not a directory", dir)
		}

		return dir, nil
	}
}

func IntValidator(input string) (interface{}, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}

	return i, nil
}
