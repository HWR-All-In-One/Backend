package main

import (
	"log"

	"github.com/HWR-All-In-One/Backend/internal/pkg/ctrl"
	_ "github.com/HWR-All-In-One/Backend/migrations"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := ctrl.New()

	migratecmd.MustRegister(app.PB, app.PB.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	err := app.PB.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	app.AddHooks()

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}

}