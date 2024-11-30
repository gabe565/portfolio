package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("deleted = false")

		// add
		new_deleted := &core.BoolField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "uenlwh8j",
			"name": "deleted",
			"type": "bool",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_deleted)
		collection.Fields.Add(new_deleted)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		// remove
		collection.Fields.RemoveById("uenlwh8j")

		return app.Save(collection)
	})
}
