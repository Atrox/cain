package store

import "path/filepath"

type Config struct {
	RetrievePath string `yaml:"defaultRetrievePath"`
	LockFile     string `yaml:"lockFile"`

	Destinations  `yaml:"destinations"`
	NamingSchemes `yaml:"namingSchemes"`
}

// Destinations locations
type Destinations struct {
	Movie  string `yaml:"movie"`
	Series string `yaml:"series"`
	Anime  string `yaml:"anime"`
	Music  string `yaml:"music"`
}

// NamingSchemes templates
type NamingSchemes struct {
	Movie  string `yaml:"movie"`
	Series string `yaml:"series"`
	Anime  string `yaml:"anime"`
	Music  string `yaml:"music"`
}

var defaultNamingSchemes = NamingSchemes{
	Movie:  "{n} ({y})/{n}",
	Series: "{n}/Season {s.pad(2)}/{n} - {s00e00} - {t}",
	Anime:  "{n}/Season {s.pad(2)}/{n} - {s00e00} - {t}",
	Music:  "{n}/{fn}",
}

func NewConfig() *Config {
	return &Config{
		LockFile:      filepath.Join(Base, "filebot.lock"),
		NamingSchemes: defaultNamingSchemes,
	}
}
