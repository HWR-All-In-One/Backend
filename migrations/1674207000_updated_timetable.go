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
		collection.Schema.RemoveField("xqx0cqlo")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("39gquw3xrrx8ul2")
		if err != nil {
			return err
		}

		// add
		del_summary := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "xqx0cqlo",
			"name": "summary",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_summary)
		collection.Schema.AddField(del_summary)

		return dao.SaveCollection(collection)
	})
}
