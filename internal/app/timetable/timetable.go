package timetable

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

type Environment struct {
	PB *pocketbase.PocketBase
}

const URL = "https://moodle.hwr-berlin.de/fb2-stundenplan/download.php?doctype=.ics&url=./fb2-stundenplaene/informatik/semester4/kursa"

func (env *Environment) List(c echo.Context) error {
	authRecord, _ := c.Get(apis.ContextAdminKey).(*models.Admin)
	fmt.Println(authRecord)
	if authRecord == nil {
		return apis.NewForbiddenError("Only auth records can access this endpoint!", nil)
	}

	return c.String(200, "Hello world!")
}
