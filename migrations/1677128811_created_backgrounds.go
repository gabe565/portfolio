package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "d6olo7t8r2b0qew",
			"created": "2023-02-23 05:06:51.422Z",
			"updated": "2023-02-23 05:06:51.422Z",
			"name": "backgrounds",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "xjslwkzi",
					"name": "url",
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
		collection, err := app.FindCollectionByNameOrId("d6olo7t8r2b0qew")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
