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

var base string

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	base = filepath.Join(home, ".config", "cain")
}

func Get(i interface{}) error {
	location := Path(i)

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
	location := Path(i)

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

func Path(i interface{}) string {
	name := strings.ToLower(reflect.TypeOf(i).Elem().Name())
	return filepath.Join(base, name+".yaml")
}
