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
		new_image_padding := &core.NumberField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "dmz2guhd",
			"name": "image_padding",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_image_padding)
		collection.Fields.Add(new_image_padding)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// remove
		collection.Fields.RemoveById("dmz2guhd")

		return app.Save(collection)
	})
}
