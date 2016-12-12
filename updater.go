package main

import (
	"fmt"
	"log"

	"github.com/atrox/cain/store"
	"github.com/equinox-io/equinox"
)

const appID = "app_4SbsY14WUcg"

// public portion of signing key generated by `equinox genkey`
var publicKey = []byte(`
-----BEGIN ECDSA PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEGj6kWvbx0d2YcEmM2KijWSnTenwsSB22
/AH1+uK39i6Jr4mfIN20WR1+w2jtxCnmqiz0yiomSJA5TxJ9mbyp/G0o3/0Rt/Yq
3F9GZBmWbduYawWwP5pgPduPWZlTDf7Q
-----END ECDSA PUBLIC KEY-----
`)

func equinoxUpdate() error {
	var opts equinox.Options
	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
		return err
	}

	// check for the update
	resp, err := equinox.Check(appID, opts)
	switch {
	case err == equinox.NotAvailableErr:
		fmt.Println("[+] No update available, already at the latest version!")
		return nil
	case err != nil:
		fmt.Println("[!] Update failed:", err)
		return err
	}

	// fetch the update and apply it
	err = resp.Apply()
	if err != nil {
		return err
	}

	fmt.Printf("[+] Updated to new version: %s!\n", resp.ReleaseVersion)
	return nil
}

var updatesChan = make(chan string)

func checkForUpdates() {
	go func(quit chan string) {
		updater := &store.Updater{}
		store.Get(updater)

		if !updater.ShouldCheck() {
			quit <- ""
			return
		}

		updater.SetNext()
		store.Save(updater)

		var opts equinox.Options
		err := opts.SetPublicKeyPEM(publicKey)
		if err != nil {
			log.Fatal(err)
		}

		_, err = equinox.Check(appID, opts)
		if err != nil {
			quit <- ""
			return
		}

		quit <- b.String("New Version of Cain is available!",
			"Update automatically with 'cain update'")
	}(updatesChan)
}
