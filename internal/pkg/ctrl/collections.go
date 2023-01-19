package ctrl

import (
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

func (a *App) CreateTimetableCollection() error {
	c := &models.Collection{}

	form := forms.NewCollectionUpsert(a.PB, c)

	form.Name = "timetable"
	form.Type = models.CollectionTypeBase
	form.Schema.AddField(&schema.SchemaField{
		Name:     "start",
		Type:     schema.FieldTypeDate,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "end",
		Type:     schema.FieldTypeDate,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "description",
		Type:     schema.FieldTypeText,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "summary",
		Type:     schema.FieldTypeText,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "location",
		Type:     schema.FieldTypeText,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "organizer",
		Type:     schema.FieldTypeText,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "type",
		Type:     schema.FieldTypeText,
		Required: true,
	})

	form.Schema.AddField(&schema.SchemaField{
		Name:     "name",
		Type:     schema.FieldTypeText,
		Required: true,
	})

	return form.Submit()

}
