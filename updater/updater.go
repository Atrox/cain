package updater

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/atrox/box"
	"github.com/atrox/cain/config"
	"github.com/tj/go-update"
	"github.com/tj/go-update/progress"
)

var (
	Version string
	b       = box.New()
)

type Updater struct {
	NextCheck time.Time `yaml:"nextCheck"`

	project         *update.Project
	automaticUpdate bool
	updateAvailable chan *update.Release
}

func New(autoUpdate bool) *Updater {
	updater := new(Updater)
	config.Storage.Get(updater)

	updater.project = &update.Project{
		Command: "cain",
		Owner:   "Atrox",
		Repo:    "cain",
		Version: Version,
	}
	updater.automaticUpdate = autoUpdate
	updater.updateAvailable = make(chan *update.Release)

	go updater.check()
	return updater
}

func (u *Updater) Run() (bool, error) {
	release := <-u.updateAvailable
	if release == nil {
		return false, nil
	}

	if !u.automaticUpdate {
		b.Println(
			fmt.Sprintf("%s (%s)", release.Version, release.PublishedAt.Format("02.01.2006")), "",
			"New Version of Cain is available!",
			"Update instantly with 'cain update'")

		return false, nil
	}

	return true, u.update(release)
}

func (u *Updater) check() {
	if !u.shouldCheck() {
		u.updateAvailable <- nil
		return
	}

	u.setNext()

	// fetch latest release
	release, err := u.fetchRelease()
	if err != nil {
		u.updateAvailable <- nil
		return
	}

	u.updateAvailable <- release
}

func (u *Updater) fetchRelease() (*update.Release, error) {
	// fetch the new releases
	releases, err := u.project.LatestReleases()
	if err != nil {
		return nil, err
	}

	// no updates
	if len(releases) == 0 {
		return nil, nil
	}

	// latest release
	latest := releases[0]
	return latest, nil
}

func (u *Updater) update(release *update.Release) error {
	b.Println("Updating...", "", "Please don't close me while I'm working")

	// find the tarball for this system
	asset := release.FindTarball(runtime.GOOS, runtime.GOARCH)
	if asset == nil {
		b.Println("No binary for your system is published", "", "Go to https://github.com/Atrox/cain/releases")
		return errors.New("No binary for your system is published on GitHub")
	}

	// download tarball to a tmp dir
	tarball, err := asset.DownloadProxy(progress.Reader)
	if err != nil {
		return err
	}

	// install it
	if err := u.project.Install(tarball); err != nil {
		return err
	}

	b.Println(fmt.Sprintf("%s (%s)", release.Version, release.PublishedAt.Format("02.01.2006")), "",
		"Successfully updated!")
	return nil
}

func (u *Updater) shouldCheck() bool {
	now := time.Now().UTC()
	return u.NextCheck.Before(now)
}

func (u *Updater) setNext() {
	now := time.Now().UTC()

	u.NextCheck = now.AddDate(0, 0, 1)
	config.Storage.Save(u)
}
