package ctrl

import (
	"github.com/HWR-All-In-One/Backend/internal/app/timetable"
	"github.com/pocketbase/pocketbase"
)

type App struct {
	PB    *pocketbase.PocketBase
	Views *Views
}

type Views struct {
	Timetable *timetable.Environment
}

func New(pb *pocketbase.PocketBase) *App {
	return &App{
		Views: &Views{
			Timetable: &timetable.Environment{
				PB: pb,
			},
		},
	}
}

func (a *App) Run() error {
	return a.PB.Start()
}
