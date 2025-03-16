package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		collection, err := app.FindCollectionByNameOrId("lf06qip3f4d11yk")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("((@request.auth.id != \"\" && trail.author = @request.auth.id) || trail.public = true) || author = @request.auth.id")

		collection.ViewRule = types.Pointer("((@request.auth.id != \"\" && trail.author = @request.auth.id) || trail.public = true) || author = @request.auth.id")

		collection.CreateRule = types.Pointer("@request.auth.id != \"\" && (trail.author = @request.auth.id || trail.public = true)")

		collection.UpdateRule = types.Pointer("@request.auth.id = author")

		collection.DeleteRule = types.Pointer("@request.auth.id = author")

		return app.Save(collection)
	}, func(app core.App) error {

		collection, err := app.FindCollectionByNameOrId("lf06qip3f4d11yk")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		collection.ViewRule = types.Pointer("")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("")

		collection.DeleteRule = types.Pointer("")

		return app.Save(collection)
	})
}
