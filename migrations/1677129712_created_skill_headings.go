package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "0tlg77ih9oej4cn",
			"created": "2023-02-23 05:21:52.018Z",
			"updated": "2023-02-23 05:21:52.018Z",
			"name": "skill_headings",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "4ksy2ptx",
					"name": "title",
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
					"id": "ciu0mpir",
					"name": "order",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null
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
		collection, err := app.FindCollectionByNameOrId("0tlg77ih9oej4cn")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
