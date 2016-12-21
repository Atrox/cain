package input

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

func PathValidator(allowEmpty bool) ValidatorFunction {
	return func(input string) (interface{}, error) {
		if input == "" {
			if allowEmpty {
				return input, nil
			}
			return nil, fmt.Errorf("This option is required")
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

func BooleanValidator(allowEmpty bool) ValidatorFunction {
	return func(input string) (interface{}, error) {
		if input == "" {
			if allowEmpty {
				return true, nil
			}
			return nil, fmt.Errorf("Please choose Yes or No")
		}

		b, err := parseBool(input)
		if err != nil {
			return nil, err
		}

		return b, nil
	}
}

// parseBool returns the boolean value represented by the string.
// Based on strconv.ParseBool but with y, yes, n, no added.
func parseBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "1", "t", "true", "y", "yes":
		return true, nil
	case "0", "f", "false", "n", "no":
		return false, nil
	}
	return false, fmt.Errorf("%s is not valid. Please choose between Yes or No", str)
}
