package store

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	yaml "gopkg.in/yaml.v2"
)

var Base string

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	Base = filepath.Join(home, ".config", "cain")
}

func Get(i interface{}) error {
	location := locationFromInterface(i)

	file, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, i)
	if err != nil {
		return err
	}

	return nil
}

func Save(i interface{}) error {
	location := locationFromInterface(i)

	b, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(location), os.ModePerm)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(location, b, 0644)
}

func GetOrCreate(i interface{}) error {
	err := Get(i)
	if err != nil {
		err := Save(i)
		if err != nil {
			return err
		}
	}

	return nil
}

func Exists(name string) bool {
	_, err := os.Stat(location(name))
	if err != nil {
		return false
	}

	return true
}

func locationFromInterface(i interface{}) string {
	return location(strings.ToLower(reflect.TypeOf(i).Elem().Name()))
}

func location(name string) string {
	return filepath.Join(Base, name+".yaml")
}
