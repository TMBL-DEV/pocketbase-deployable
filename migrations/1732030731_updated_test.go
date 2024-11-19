package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tifhgkbl0dpf6rl")
		if err != nil {
			return err
		}

		collection.Name = "blogs"

		// update
		edit_richt_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jhvyndlw",
			"name": "richt_content",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), edit_richt_content); err != nil {
			return err
		}
		collection.Schema.AddField(edit_richt_content)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tifhgkbl0dpf6rl")
		if err != nil {
			return err
		}

		collection.Name = "test"

		// update
		edit_richt_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jhvyndlw",
			"name": "field",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), edit_richt_content); err != nil {
			return err
		}
		collection.Schema.AddField(edit_richt_content)

		return dao.SaveCollection(collection)
	})
}
