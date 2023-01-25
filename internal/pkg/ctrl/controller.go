package ctrl

import (
	"github.com/HWR-All-In-One/Backend/internal/pkg/timetable"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

const URL = "https://moodle.hwr-berlin.de/fb2-stundenplan/download.php?doctype=.ics&url=./fb2-stundenplaene/informatik/semester4/kursa"

type App struct {
	PB *pocketbase.PocketBase
}

func New(pb *pocketbase.PocketBase) *App {
	return &App{
		PB: pb,
	}
}

func (a *App) Run() error {
	return a.PB.Start()
}

func (a *App) InsertTimetableData() error {
	tt := timetable.New(URL)

	err := tt.Parse()

	if err != nil {
		return err
	}

	collection, err := a.PB.Dao().FindCollectionByNameOrId("timetable")

	if err != nil {
		return err
	}

	for _, lesson := range tt.Lessons {
		record := models.NewRecord(collection)
		record.Set("start", lesson.Start)
		record.Set("end", lesson.End)
		record.Set("location", lesson.Room)
		record.Set("organizer", lesson.Teacher)
		record.Set("type", lesson.Kind)
		record.Set("name", lesson.Name)

		err := a.PB.Dao().SaveRecord(record)

		if err != nil {
			return err
		}

	}

	return nil
}
