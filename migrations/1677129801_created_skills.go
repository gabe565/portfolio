package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "4wkecakvi7gdg80",
			"created": "2023-02-23 05:23:21.017Z",
			"updated": "2023-02-23 05:23:21.017Z",
			"name": "skills",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "7odfbq6r",
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
					"id": "qmvm1gcf",
					"name": "heading",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "0tlg77ih9oej4cn",
						"cascadeDelete": false,
						"maxSelect": 1,
						"displayFields": [
							"title"
						]
					}
				},
				{
					"system": false,
					"id": "erys1qxl",
					"name": "rating",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 5
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
		collection, err := app.FindCollectionByNameOrId("4wkecakvi7gdg80")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
