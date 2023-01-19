package ctrl

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

type App struct {
	PB          *pocketbase.PocketBase
	Collections []*models.Collection
}

func New(pb *pocketbase.PocketBase) *App {
	return &App{
		PB: pb,
	}
}

func (a *App) Run() error {
	return a.PB.Start()
}
