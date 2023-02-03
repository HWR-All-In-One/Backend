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
		edit_hwr_password := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qsefvlxh",
			"name": "hwr_password",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_hwr_password)
		collection.Schema.AddField(edit_hwr_password)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update
		edit_hwr_password := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qsefvlxh",
			"name": "hwr_password",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 16,
				"max": null,
				"pattern": ""
			}
		}`), edit_hwr_password)
		collection.Schema.AddField(edit_hwr_password)

		return dao.SaveCollection(collection)
	})
}
