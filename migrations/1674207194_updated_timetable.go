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

		// remove
		collection.Schema.RemoveField("vlqpixt2")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("39gquw3xrrx8ul2")
		if err != nil {
			return err
		}

		// add
		del_description := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "vlqpixt2",
			"name": "description",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_description)
		collection.Schema.AddField(del_description)

		return dao.SaveCollection(collection)
	})
}
