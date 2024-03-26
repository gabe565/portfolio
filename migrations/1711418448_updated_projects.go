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

		collection, err := dao.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// add
		new_priority := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qyhpws8u",
			"name": "priority",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": true
			}
		}`), new_priority); err != nil {
			return err
		}
		collection.Schema.AddField(new_priority)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qyhpws8u")

		return dao.SaveCollection(collection)
	})
}
