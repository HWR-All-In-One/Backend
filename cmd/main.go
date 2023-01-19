package main

import (
	"log"

	"github.com/HWR-All-In-One/Backend/internal/pkg/ctrl"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := ctrl.App{
		PB: pocketbase.New(),
	}

	migratecmd.MustRegister(app.PB, app.PB.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	err := app.PB.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	app.InsertTimetableData()

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}

}
