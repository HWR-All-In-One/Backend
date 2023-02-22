migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("_pb_users_auth_")

  // remove
  collection.schema.removeField("vfd8l3dk")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("_pb_users_auth_")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vfd8l3dk",
    "name": "passwort",
    "type": "text",
    "required": true,
    "unique": false,
    "options": {
      "min": 10,
      "max": 25,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
})
