migrate((db) => {
  const collection = new Collection({
    "id": "9vgfsjl6k0gih7n",
    "created": "2023-01-17 13:42:35.769Z",
    "updated": "2023-01-17 13:42:35.769Z",
    "name": "timetable",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "zqsm2gxn",
        "name": "start",
        "type": "date",
        "required": true,
        "unique": false,
        "options": {
          "min": "",
          "max": ""
        }
      },
      {
        "system": false,
        "id": "k2vrvxws",
        "name": "end",
        "type": "date",
        "required": true,
        "unique": false,
        "options": {
          "min": "",
          "max": ""
        }
      },
      {
        "system": false,
        "id": "nbwgdy2k",
        "name": "description",
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
        "id": "w2umsdnu",
        "name": "summary",
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
        "id": "1hhholhr",
        "name": "location",
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
        "id": "1yor2fhs",
        "name": "organizer",
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
        "id": "wrp3jkqa",
        "name": "type",
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
        "id": "tmlcvkey",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("9vgfsjl6k0gih7n");

  return dao.deleteCollection(collection);
})
