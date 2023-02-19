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

		// remove
		collection.Schema.RemoveField("rnk1oaar")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add
		del_private_email := &schema.SchemaField{}
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
		}`), del_private_email)
		collection.Schema.AddField(del_private_email)

		return dao.SaveCollection(collection)
	})
}
