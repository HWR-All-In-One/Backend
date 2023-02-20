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
		collection.Schema.RemoveField("nkotkpy8")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add
		del_hwr_email := &schema.SchemaField{}
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
		}`), del_hwr_email)
		collection.Schema.AddField(del_hwr_email)

		return dao.SaveCollection(collection)
	})
}
