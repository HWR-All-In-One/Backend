package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "rfxs4rahsiaxc9a",
			"created": "2023-02-17 09:02:06.334Z",
			"updated": "2023-02-17 09:02:06.334Z",
			"name": "timetable_updates",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "iesvnxwc",
					"name": "profession",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "jvzhfrfj",
					"name": "semester",
					"type": "number",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null
					}
				},
				{
					"system": false,
					"id": "eyqec78l",
					"name": "group",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": 1,
						"pattern": "^[a-z]$"
					}
				},
				{
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
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rfxs4rahsiaxc9a")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
