package filebot

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/atrox/cain/store"
)

type FileBot struct {
	RetrievePath      string
	NonStrictMatching bool

	executable string
	config     *store.Config
	args       *args
}

func New(conf *store.Config) (*FileBot, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}

	return &FileBot{
		executable: path,
		config:     conf,
		args:       newArgs(),
	}, nil
}

func (f *FileBot) Execute() error {
	retrievePath := f.config.DefaultRetrievePath

	// if path flag specified, overwrite retrievePath
	if f.RetrievePath != "" {
		retrievePath = f.RetrievePath
	}

	if retrievePath == "" {
		return fmt.Errorf("[!] '--path' or 'defaultRetrievePath' not specified")
	}

	if f.config.Language != "" {
		f.args.Add("--lang", f.config.Language)
	}

	if f.config.LogFile != "" {
		f.args.Add("--log-file", filepath.Clean(f.config.LogFile))
	}

	if f.config.CleanupAfterwards {
		f.args.AddDefinition("clean", "y")
	}

	if f.NonStrictMatching || f.config.NonStrictMatching {
		f.args.Add("-non-strict")
	}

	f.addPaths()
	f.addNotifiers()

	f.args.Add(filepath.Clean(retrievePath))

	cmd := exec.Command(f.executable, *f.args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

var paths = []string{"excludeList", "movieFormat", "seriesFormat", "animeFormat", "musicFormat"}

func (f *FileBot) addPaths() {
	for _, name := range paths {
		f.args.AddDefinition(name, f.getPath(name))
	}
}

func (f *FileBot) addNotifiers() {
	typ := reflect.TypeOf(&f.config.Notifiers).Elem()
	value := reflect.ValueOf(&f.config.Notifiers).Elem()

	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		field := value.Field(i)

		switch in := field.Interface().(type) {
		case string:
			f.args.AddDefinition(strings.ToLower(name), in)
		}
	}
}

func (f *FileBot) getPath(name string) string {
	switch name {
	case "excludeList":
		if f.config.LockFile == "" {
			return ""
		}

		return filepath.Clean(f.config.LockFile)
	case "movieFormat":
		if f.config.Destinations.Movie == "" {
			return ""
		}

		return filepath.Join(f.config.Destinations.Movie, f.config.NamingSchemes.Movie)
	case "seriesFormat":
		if f.config.Destinations.Series == "" {
			return ""
		}

		return filepath.Join(f.config.Destinations.Series, f.config.NamingSchemes.Series)
	case "animeFormat":
		if f.config.Destinations.Anime == "" {
			return ""
		}

		return filepath.Join(f.config.Destinations.Anime, f.config.NamingSchemes.Anime)
	case "musicFormat":
		if f.config.Destinations.Music == "" {
			return ""
		}

		f.args.AddDefinition("music", "y")
		return filepath.Join(f.config.Destinations.Music, f.config.NamingSchemes.Music)
	default:
		return ""
	}
}
