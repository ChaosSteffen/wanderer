package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"id": "4wbv9tz5zjdrjh1",
			"created": "2024-04-01 17:47:10.015Z",
			"updated": "2024-04-01 17:47:10.015Z",
			"name": "trails_filter",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ovledory",
					"name": "max_distance",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 1
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT (ROW_NUMBER() OVER()) as id, MAX(trails.distance) AS max_distance FROM trails;"
			}
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {

		collection, err := app.FindCollectionByNameOrId("4wbv9tz5zjdrjh1")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
