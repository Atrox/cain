package updater

import (
	"fmt"
	"log"
	"time"

	"github.com/atrox/box"
	"github.com/atrox/cain/config"
	"github.com/equinox-io/equinox"
)

var b = box.New()

type Updater struct {
	NextCheck time.Time `yaml:"nextCheck"`

	automaticUpdate bool
	updateAvailable chan *equinox.Response
}

func New(autoUpdate bool) *Updater {
	updater := new(Updater)
	config.Storage.Get(updater)

	updater.automaticUpdate = autoUpdate
	updater.updateAvailable = make(chan *equinox.Response)

	go updater.check()
	return updater
}

func (u *Updater) Run() error {
	update := <-u.updateAvailable
	if update == nil {
		return nil
	}

	if !u.automaticUpdate {
		b.Println(
			fmt.Sprintf("%s (%s)", update.ReleaseTitle, update.ReleaseDate.Format("02.01.2006")), "",
			"New Version of Cain is available!",
			"Update automatically with 'cain update'")

		return nil
	}

	b.Println("Updating...", "", "Please don't close me while I'm working")

	// fetch the update and apply it
	err := update.Apply()
	if err != nil {
		return err
	}

	b.Println(fmt.Sprintf("%s (%s)", update.ReleaseTitle, update.ReleaseDate.Format("02.01.2006")), "",
		"Successfully updated!")
	return nil
}

func ForceRun() error {
	var opts equinox.Options
	err := opts.SetPublicKeyPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	update, err := equinox.Check(appID, opts)
	if err != nil {
		return err
	}

	b.Println("Updating...", "", "Please don't close me while I'm working")

	// fetch the update and apply it
	err = update.Apply()
	if err != nil {
		return err
	}

	b.Println(fmt.Sprintf("%s (%s)", update.ReleaseTitle, update.ReleaseDate.Format("02.01.2006")), "",
		"Successfully updated!")
	return nil
}

func (u *Updater) check() {
	if !u.shouldCheck() {
		u.updateAvailable <- nil
		return
	}

	u.setNext()

	var opts equinox.Options
	err := opts.SetPublicKeyPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := equinox.Check(appID, opts)
	if err != nil {
		u.updateAvailable <- nil
		return
	}

	u.updateAvailable <- &resp
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
