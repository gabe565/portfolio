package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// add
		new_priority := &core.NumberField{}
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
		collection.Fields.Add(new_priority)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// remove
		collection.Fields.RemoveById("qyhpws8u")

		return app.Save(collection)
	})
}
