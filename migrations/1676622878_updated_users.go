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

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update
		edit_group := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "j7nkw2oq",
			"name": "group",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": 1,
				"pattern": "^\\w$"
			}
		}`), edit_group)
		collection.Schema.AddField(edit_group)

		// update
		edit_semester := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "38soy3h9",
			"name": "semester",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), edit_semester)
		collection.Schema.AddField(edit_semester)

		// update
		edit_profession := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rhwkhqut",
			"name": "profession",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_profession)
		collection.Schema.AddField(edit_profession)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update
		edit_group := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "j7nkw2oq",
			"name": "course_title",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": "IT[0-9]{2}[A-Z]{0,1}"
			}
		}`), edit_group)
		collection.Schema.AddField(edit_group)

		// update
		edit_semester := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "38soy3h9",
			"name": "year",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), edit_semester)
		collection.Schema.AddField(edit_semester)

		// update
		edit_profession := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rhwkhqut",
			"name": "semester",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_profession)
		collection.Schema.AddField(edit_profession)

		return dao.SaveCollection(collection)
	})
}
