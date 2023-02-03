package ctrl

import (
	"github.com/HWR-All-In-One/Backend/internal/app/timetable"
	"github.com/HWR-All-In-One/Backend/internal/pkg/safe"
	"github.com/pocketbase/pocketbase"
)

type App struct {
	PB    *pocketbase.PocketBase
	Views *Views
	Safe  safe.Getter
}

type Views struct {
	Timetable *timetable.Environment
}

func New() *App {
	pb := pocketbase.New()
	return &App{
		PB:   pb,
		Safe: safe.New(),
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
