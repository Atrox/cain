package config

import (
	"log"

	"github.com/atrox/store"
)

var Storage *store.Store

func init() {
	var err error
	Storage, err = store.New("cain")
	if err != nil {
		log.Fatal(err)
	}
}
