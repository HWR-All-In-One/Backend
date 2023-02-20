package hwr

import (
	"errors"

	aes "github.com/HWR-All-In-One/Backend/internal/pkg/aes"
	"github.com/HWR-All-In-One/Backend/internal/pkg/hwr"
	"github.com/HWR-All-In-One/Backend/internal/pkg/safe"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type Environment struct {
	PB   *pocketbase.PocketBase
	Safe *safe.Safe
}

func (env *Environment) ValidateUser(e *core.RecordCreateEvent) error {
	if e.Record.Collection().Name == "users" {
		key := env.Safe.Get()
		password := e.Record.GetString("hwr_password")
		email := e.Record.GetString("email")

		isValid, err := hwr.ValidateUser(email, password)

		if err != nil {
			return err
		}

		if !isValid {
			return errors.New("hwr user does not exist")
		}

		enc, err := aes.Encrypt(key, password)

		if err != nil {
			return err
		}

		e.Record.Set("hwr_password", enc)

		// checking if the user has the timetable ready for him
		collection, err := env.PB.Dao().FindCollectionByNameOrId("timetable")

		if err != nil {
			return err
		}

		query := env.PB.Dao().RecordQuery(collection).AndWhere(dbx.HashExp{
			"profession": e.Record.Get("profession"),
			"semester":   e.Record.Get("semester"),
			"group":      e.Record.Get("group"),
		})

		row := dbx.NullStringMap{}
		err = query.One(&row)

		if err == nil {
			return nil
		}

		// get the data and insert it into the database

		return nil
	}
	return nil
}
