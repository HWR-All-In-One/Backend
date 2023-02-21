package timetable

import (
	"errors"
	"time"

	"github.com/HWR-All-In-One/Backend/internal/pkg/timetable"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type Environment struct {
	PB *pocketbase.PocketBase
}

// CheckStatus will check the update status of the records and if the records
// for the given profession, semester and group are not available then it will be
// created
func (env *Environment) CheckStatus(e *core.RecordsListEvent) error {

	authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
	if authRecord == nil {
		return errors.New("you are not loged in")
	}

	profession := authRecord.GetString("profession")
	group := authRecord.GetString("group")
	semester := authRecord.GetString("semester")

	tt, err := timetable.New(profession, semester, group)

	if err != nil {
		return err
	}

	collection, err := env.PB.Dao().FindCollectionByNameOrId("timetable")

	if err != nil {
		return err
	}

	query := env.PB.Dao().RecordQuery(collection).AndWhere(dbx.HashExp{
		"profession": profession,
		"group":      group,
		"semester":   semester,
	})

	row := dbx.NullStringMap{}

	if err := query.One(&row); err != nil {
		return err
	}

	record := models.NewRecordFromNullStringMap(collection, row)

	if record != nil {
		updated := record.GetTime("updated")
		now := time.Now()

		_ = updated
		_ = now
		// records exits, check if they are old, if not serve if yes fetch new records,
		// replace the old ones wiht the new ones keep the history.
		return nil
	}

	// fetch the records new becuase they do not exist
	err = tt.GetNewRecords(collection)

	if err != nil {
		return err
	}

	err = env.PB.Dao().RunInTransaction(func(txDao *daos.Dao) error {
		for _, record := range tt.Records {
			err := txDao.SaveRecord(record)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
