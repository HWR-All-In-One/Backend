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

		// add
		new_hwr_email := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "nkotkpy8",
			"name": "hwr_email",
			"type": "email",
			"required": true,
			"unique": true,
			"options": {
				"exceptDomains": null,
				"onlyDomains": [
					"stud.hwr-berlin.de",
					"hwr-berlin.de"
				]
			}
		}`), new_hwr_email)
		collection.Schema.AddField(new_hwr_email)

		// add
		new_hwr_password := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qsefvlxh",
			"name": "hwr_password",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 10,
				"max": 25,
				"pattern": ""
			}
		}`), new_hwr_password)
		collection.Schema.AddField(new_hwr_password)

		// add
		new_course_title := &schema.SchemaField{}
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
		}`), new_course_title)
		collection.Schema.AddField(new_course_title)

		// add
		new_year := &schema.SchemaField{}
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
		}`), new_year)
		collection.Schema.AddField(new_year)

		// add
		new_course := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rhwkhqut",
			"name": "course",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_course)
		collection.Schema.AddField(new_course)

		// add
		new_private_email := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rnk1oaar",
			"name": "private_email",
			"type": "email",
			"required": true,
			"unique": true,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), new_private_email)
		collection.Schema.AddField(new_private_email)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("nkotkpy8")

		// remove
		collection.Schema.RemoveField("qsefvlxh")

		// remove
		collection.Schema.RemoveField("j7nkw2oq")

		// remove
		collection.Schema.RemoveField("38soy3h9")

		// remove
		collection.Schema.RemoveField("rhwkhqut")

		// remove
		collection.Schema.RemoveField("rnk1oaar")

		return dao.SaveCollection(collection)
	})
}
