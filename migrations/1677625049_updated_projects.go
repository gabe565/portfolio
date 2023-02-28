package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("deleted = false")

		// add
		new_deleted := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "uenlwh8j",
			"name": "deleted",
			"type": "bool",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_deleted)
		collection.Schema.AddField(new_deleted)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		// remove
		collection.Schema.RemoveField("uenlwh8j")

		return dao.SaveCollection(collection)
	})
}
