package timetable

import (
	"fmt"
	"net/http"
	"net/url"

	ics "github.com/arran4/golang-ical"
	"github.com/pocketbase/pocketbase/models"
)

const description = 6

type Timetable struct {
	Profession string
	Semester   string
	Group      string
	Url        *url.URL
	Calendar   *ics.Calendar
	Records    []*models.Record
}

func New(profession, semester, group string) (*Timetable, error) {
	tt := &Timetable{
		Profession: profession,
		Semester:   semester,
		Group:      group,
	}

	return tt, tt.createUrl()
}

func (t *Timetable) createUrl() error {
	semester := "semester" + t.Semester
	group := "kurs" + t.Group
	path, err := url.JoinPath("fb2-stundenplaene", t.Profession, semester, group)

	if err != nil {
		return err
	}

	icsUrl := &url.URL{
		Scheme: "https",
		Host:   "moodle.hwr-berlin.de",
		Path:   "fb2-stundenplan/download.php",
	}

	values := icsUrl.Query()
	values.Add("doctype", ".ics")
	values.Add("url", "./"+path)
	decoded, err := url.QueryUnescape(values.Encode())

	if err != nil {
		return err
	}

	icsUrl.RawQuery = decoded

	t.Url = icsUrl

	return nil
}

func (t *Timetable) parse() error {
	resp, err := http.Get(t.Url.String())

	if err != nil {
		return err
	}

	cal, err := ics.ParseCalendar(resp.Body)

	if err != nil {
		return err
	}

	t.Calendar = cal

	return nil
}

func (t *Timetable) GetNewRecords(collection *models.Collection) error {
	err := t.parse()

	if err != nil {
		return err
	}

	for _, lesson := range t.Calendar.Events() {
		record := models.NewRecord(collection)
		rawDesc := lesson.Properties[description].Value
		desc := decodeDescription(rawDesc)
		start, err := lesson.GetStartAt()
		if err != nil {
			return err
		}
		end, err := lesson.GetEndAt()
		if err != nil {
			return err
		}
		record.Set("start", start)
		record.Set("end", end)
		t.setAllFields(record, desc)
		t.Records = append(t.Records, record)
	}
	return nil
}
