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

		collection, err := dao.FindCollectionByNameOrId("rfxs4rahsiaxc9a")
		if err != nil {
			return err
		}

		// update
		edit_last_query := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "8ckoubpi",
			"name": "last_query",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_last_query)
		collection.Schema.AddField(edit_last_query)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rfxs4rahsiaxc9a")
		if err != nil {
			return err
		}

		// update
		edit_last_query := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "8ckoubpi",
			"name": "last_update",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_last_query)
		collection.Schema.AddField(edit_last_query)

		return dao.SaveCollection(collection)
	})
}
