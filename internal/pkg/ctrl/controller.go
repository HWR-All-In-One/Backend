package ctrl

import (
	"github.com/HWR-All-In-One/Backend/internal/app/hwr"
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
	HWR       *hwr.Environment
}

func New() *App {
	pb := pocketbase.New()
	safe := safe.New()
	return &App{
		PB:   pb,
		Safe: safe,
		Views: &Views{
			Timetable: &timetable.Environment{
				PB: pb,
			},
			HWR: &hwr.Environment{
				PB:   pb,
				Safe: safe,
			},
		},
	}
}

func (a *App) Run() error {
	return a.PB.Start()
}
