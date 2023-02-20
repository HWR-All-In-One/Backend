package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("39gquw3xrrx8ul2")
		if err != nil {
			return err
		}

		// add
		new_pause := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "8kmsocrf",
			"name": "pause",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": 0,
				"max": null
			}
		}`), new_pause)
		collection.Schema.AddField(new_pause)

		// add
		new_notice := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "o83rzcwp",
			"name": "notice",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_notice)
		collection.Schema.AddField(new_notice)

		// update
		edit_room := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "m82nwpcf",
			"name": "room",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_room)
		collection.Schema.AddField(edit_room)

		// update
		edit_teacher := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fc79xyvz",
			"name": "teacher",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_teacher)
		collection.Schema.AddField(edit_teacher)

		// update
		edit_kind := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rxlng4vt",
			"name": "kind",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_kind)
		collection.Schema.AddField(edit_kind)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("39gquw3xrrx8ul2")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("8kmsocrf")

		// remove
		collection.Schema.RemoveField("o83rzcwp")

		// update
		edit_room := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "m82nwpcf",
			"name": "location",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_room)
		collection.Schema.AddField(edit_room)

		// update
		edit_teacher := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "fc79xyvz",
			"name": "organizer",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_teacher)
		collection.Schema.AddField(edit_teacher)

		// update
		edit_kind := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rxlng4vt",
			"name": "type",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_kind)
		collection.Schema.AddField(edit_kind)

		return dao.SaveCollection(collection)
	})
}
