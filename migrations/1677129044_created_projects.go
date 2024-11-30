package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "iowdsbt17krdtf0",
			"created": "2023-02-23 05:10:44.112Z",
			"updated": "2023-02-23 05:10:44.112Z",
			"name": "projects",
			"type": "base",
			"system": false,
			"schema": [
				{
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
				},
				{
					"system": false,
					"id": "7mmshm0g",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "15zk9ggm",
					"name": "url",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "ibl7dxal",
					"name": "description",
					"type": "editor",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "wqlxu0nh",
					"name": "tags",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "a3ukxinbygb6okj",
						"cascadeDelete": false,
						"maxSelect": null,
						"displayFields": [
							"title"
						]
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
		collection, err := app.FindCollectionByNameOrId("iowdsbt17krdtf0")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
