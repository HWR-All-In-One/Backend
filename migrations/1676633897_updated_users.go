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
		new_timetable_update := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "tjt3d1xn",
			"name": "timetable_update",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "rfxs4rahsiaxc9a",
				"cascadeDelete": false
			}
		}`), new_timetable_update)
		collection.Schema.AddField(new_timetable_update)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("tjt3d1xn")

		return dao.SaveCollection(collection)
	})
}
