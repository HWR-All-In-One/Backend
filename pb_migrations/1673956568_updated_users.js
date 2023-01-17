migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("_pb_users_auth_")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "9neaixag",
    "name": "hwr_passwort",
    "type": "text",
    "required": true,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "rwsltnx8",
    "name": "kursbezeichnung",
    "type": "text",
    "required": true,
    "unique": false,
    "options": {
      "min": 4,
      "max": null,
      "pattern": "IT[0-9]{2}[A-Z]{0,1}"
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "bbozbpox",
    "name": "jahrgang",
    "type": "number",
    "required": true,
    "unique": false,
    "options": {
      "min": null,
      "max": null
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "97iauy4y",
    "name": "kurs",
    "type": "text",
    "required": true,
    "unique": false,
    "options": {
      "min": 1,
      "max": 1,
      "pattern": "[A-Z]"
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "gxoeuush",
    "name": "private_email",
    "type": "email",
    "required": true,
    "unique": true,
    "options": {
      "exceptDomains": null,
      "onlyDomains": null
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "jemooyfe",
    "name": "ist_kursprecher",
    "type": "bool",
    "required": false,
    "unique": false,
    "options": {}
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("_pb_users_auth_")

  // remove
  collection.schema.removeField("9neaixag")

  // remove
  collection.schema.removeField("rwsltnx8")

  // remove
  collection.schema.removeField("bbozbpox")

  // remove
  collection.schema.removeField("97iauy4y")

  // remove
  collection.schema.removeField("gxoeuush")

  // remove
  collection.schema.removeField("jemooyfe")

  return dao.saveCollection(collection)
})
