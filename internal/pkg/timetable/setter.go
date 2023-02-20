package timetable

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/spf13/cast"
)

func (t *Timetable) setAllFields(record *models.Record, desc map[string]string) {
	record.Set("room", desc["raum"])
	record.Set("teacher", desc["dozent"])
	record.Set("kind", desc["art"])
	record.Set("notice", desc["anmerkung"])
	record.Set("name", desc["veranstaltung"])
	record.Set("profession", t.Profession)
	record.Set("semester", t.Semester)
	record.Set("group", t.Group)
	record.Set("pause", cast.ToInt(desc["pause"]))
}
