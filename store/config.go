package store

import "path/filepath"

type Config struct {
	DefaultRetrievePath string `yaml:"defaultRetrievePath"`
	LockFile            string `yaml:"lockFile"`
	LogFile             string `yaml:"logFile"`
	AutoUpdate          bool   `yaml:"autoUpdate"`
	Language            string `yaml:"language"`
	NonStrictMatching   bool   `yaml:"nonStrictMatching"`
	CleanupAfterwards   bool   `yaml:"cleanupAfterwards"`

	Destinations  Destinations  `yaml:"destinations"`
	NamingSchemes NamingSchemes `yaml:"namingSchemes"`
	Notifiers     Notifiers     `yaml:"notifiers"`
}

// Destinations locations
type Destinations struct {
	Movie  string `yaml:"movie"`
	Series string `yaml:"series"`
	Anime  string `yaml:"anime"`
}

// NamingSchemes templates
type NamingSchemes struct {
	Movie  string `yaml:"movie"`
	Series string `yaml:"series"`
	Anime  string `yaml:"anime"`
}

// Notifiers configuration
type Notifiers struct {
	Plex       string `yaml:"plex"`
	Kodi       string `yaml:"kodi"`
	Emby       string `yaml:"emby"`
	Pushover   string `yaml:"pushover"`
	PushBullet string `yaml:"pushBullet"`
	Gmail      string `yaml:"gmail"`
	Mail       string `yaml:"mail"`
}

var defaultNamingSchemes = NamingSchemes{
	Movie:  "{n} ({y})/{n}",
	Series: "{n}/Season {s.pad(2)}/{n} - {s00e00} - {t}",
	Anime:  "{n}/Season {s.pad(2)}/{n} - {s00e00} - {t}",
}

func NewConfig() *Config {
	return &Config{
		LockFile:      filepath.Join(base, "filebot.lock"),
		LogFile:       filepath.Join(base, "filebot.log"),
		Language:      "en",
		AutoUpdate:    true,
		NamingSchemes: defaultNamingSchemes,
	}
}
