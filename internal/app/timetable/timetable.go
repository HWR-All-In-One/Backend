package timetable

import (
	"net/http"

	"github.com/HWR-All-In-One/Backend/internal/pkg/timetable"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

type Environment struct {
	PB *pocketbase.PocketBase
}

const URL = "https://moodle.hwr-berlin.de/fb2-stundenplan/download.php?doctype=.ics&url=./fb2-stundenplaene/informatik/semester4/kursa"

func (env *Environment) List(c echo.Context) error {
	tt, err := timetable.Parse(URL)

	if err != nil {
		return err
	}

	lessons, err := timetable.DecodeLessons(tt)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, lessons)
}
