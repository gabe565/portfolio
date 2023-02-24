package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
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

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("a3ukxinbygb6okj")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
