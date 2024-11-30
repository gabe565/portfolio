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
		collection.Fields.RemoveById("kgl3tdra")

		// add
		new_image := &core.FileField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "9wstjysn",
			"name": "image",
			"type": "file",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"maxSize": 1048576,
				"mimeTypes": [
					"image/jpeg",
					"image/png",
					"image/svg+xml",
					"image/gif",
					"image/webp"
				],
				"thumbs": []
			}
		}`), new_image)
		collection.Fields.Add(new_image)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// add
		del_image_path := &core.TextField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kgl3tdra",
			"name": "image_path",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_image_path)
		collection.Fields.Add(del_image_path)

		// remove
		collection.Fields.RemoveById("9wstjysn")

		return app.Save(collection)
	})
}
