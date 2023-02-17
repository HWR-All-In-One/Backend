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
				"pattern": "^[a-z]$"
			}
		}`), edit_group)
		collection.Schema.AddField(edit_group)

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

		return dao.SaveCollection(collection)
	})
}
