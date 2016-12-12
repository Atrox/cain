package store

import "time"

type Updater struct {
	NextCheck time.Time `yaml:"nextCheck"`
}

func (u *Updater) ShouldCheck() bool {
	now := time.Now().UTC()
	return u.NextCheck.Before(now)
}

func (u *Updater) SetNext() {
	now := time.Now().UTC()
	u.NextCheck = now.AddDate(0, 0, 1)
}
