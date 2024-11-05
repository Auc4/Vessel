package config

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func ConnectDB() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
