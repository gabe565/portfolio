package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "a3ukxinbygb6okj",
			"created": "2023-02-23 05:08:57.692Z",
			"updated": "2023-02-23 05:08:57.692Z",
			"name": "project_tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ieuhtmnz",
					"name": "title",
					"type": "text",
					"required": true,
					"unique": true,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("a3ukxinbygb6okj")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
