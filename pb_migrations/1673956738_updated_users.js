migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("_pb_users_auth_")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "jvlxpblj",
    "name": "hwr_email",
    "type": "email",
    "required": true,
    "unique": true,
    "options": {
      "exceptDomains": null,
      "onlyDomains": [
        "stud.hwr-berlin.de"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("_pb_users_auth_")

  // remove
  collection.schema.removeField("jvlxpblj")

  return dao.saveCollection(collection)
})
