package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "r8v4u6efmdnexqh",
			"created": "2023-02-24 07:23:28.687Z",
			"updated": "2023-02-24 07:23:28.687Z",
			"name": "contact_form",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "8t2ftyna",
					"name": "name",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "yphchlsv",
					"name": "email",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "grb4tz2g",
					"name": "message",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": "",
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
		collection, err := app.FindCollectionByNameOrId("r8v4u6efmdnexqh")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
