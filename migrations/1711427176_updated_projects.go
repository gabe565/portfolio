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

		// remove
		collection.Fields.RemoveById("dmz2guhd")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// add
		del_image_padding := &core.NumberField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "dmz2guhd",
			"name": "image_padding",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), del_image_padding); err != nil {
			return err
		}
		collection.Fields.Add(del_image_padding)

		return app.Save(collection)
	})
}
