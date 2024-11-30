package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "1ki6exglejuj8un",
			"created": "2023-02-23 05:11:21.505Z",
			"updated": "2023-02-23 05:11:21.506Z",
			"name": "redirects",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "nerjjwar",
					"name": "handle",
					"type": "text",
					"required": true,
					"unique": true,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "odmwv9hh",
					"name": "url",
					"type": "text",
					"required": true,
					"unique": false,
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
		collection, err := app.FindCollectionByNameOrId("1ki6exglejuj8un")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
