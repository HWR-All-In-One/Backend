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
		new_profession := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zsyxz1ec",
			"name": "profession",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": "^[a-z]+$"
			}
		}`), new_profession)
		collection.Schema.AddField(new_profession)

		// add
		new_semester := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "u4pirxu3",
			"name": "semester",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_semester)
		collection.Schema.AddField(new_semester)

		// add
		new_group := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "cszhqi0v",
			"name": "group",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 1,
				"max": 1,
				"pattern": "^[a-z]$"
			}
		}`), new_group)
		collection.Schema.AddField(new_group)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("39gquw3xrrx8ul2")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("zsyxz1ec")

		// remove
		collection.Schema.RemoveField("u4pirxu3")

		// remove
		collection.Schema.RemoveField("cszhqi0v")

		return dao.SaveCollection(collection)
	})
}
