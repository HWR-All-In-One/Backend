package timetable

import (
	"net/http"
	"net/url"

	"github.com/HWR-All-In-One/Backend/internal/pkg/timetable"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

type Environment struct {
	PB *pocketbase.PocketBase
}

func timetableURL(params echo.PathParams) (*url.URL, error) {
	urlValue, err := url.JoinPath("./fb2-stundeplaene", params.Get("fachrichtung", ""), params.Get("kurs", ""), params.Get("semester", ""))

	if err != nil {
		return nil, err
	}

	hwrIcsUrl := url.URL{
		Scheme: "https",
		Host:   "moodle.hwr-berlin.de",
		Path:   "/fb2-stundenplan/download.php?doctype=.ics&url=" + urlValue,
	}

	return &hwrIcsUrl, nil

}

func (env *Environment) List(c echo.Context) error {
	params := c.PathParams()
	url, err := timetableURL(params)

	if err != nil {
		return err
	}

	tt, err := timetable.Parse(url.String())

	if err != nil {
		return err
	}

	lessons, err := timetable.DecodeLessons(tt)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, lessons)
}
