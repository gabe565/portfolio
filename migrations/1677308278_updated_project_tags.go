package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("a3ukxinbygb6okj")
		if err != nil {
			return err
		}

		// add
		new_color := &core.TextField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qmk7gf66",
			"name": "color",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_color)
		collection.Fields.Add(new_color)

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("a3ukxinbygb6okj")
		if err != nil {
			return err
		}

		// remove
		collection.Fields.RemoveById("qmk7gf66")

		return app.Save(collection)
	})
}
