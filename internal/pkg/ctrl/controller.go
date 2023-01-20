package ctrl

import (
	"github.com/HWR-All-In-One/Backend/internal/pkg/timetable"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

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
	lessons, err := timetable.Parse()

	if err != nil {
		return err
	}

	collection, err := a.PB.Dao().FindCollectionByNameOrId("timetable")

	if err != nil {
		return err
	}

	for _, lesson := range lessons {
		record := models.NewRecord(collection)
		record.Set("start", lesson.Start)
		record.Set("end", lesson.End)
		record.Set("description", lesson.Description)
		record.Set("summary", lesson.Summary)
		record.Set("location", lesson.Location)
		record.Set("organizer", lesson.Organizer)
		record.Set("type", lesson.Type)
		record.Set("name", lesson.Name)

		err := a.PB.Dao().SaveRecord(record)

		if err != nil {
			return err
		}

	}

	return nil
}
