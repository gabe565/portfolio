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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kgl3tdra")

		// add
		new_image := &schema.SchemaField{}
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
		collection.Schema.AddField(new_image)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		// add
		del_image_path := &schema.SchemaField{}
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
		collection.Schema.AddField(del_image_path)

		// remove
		collection.Schema.RemoveField("9wstjysn")

		return dao.SaveCollection(collection)
	})
}
