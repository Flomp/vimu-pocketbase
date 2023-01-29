package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func findTotalFilesByOwner(dao *daos.Dao, collection *models.Collection, ownerId string) (int, error) {
	var total int

	err := dao.RecordQuery(collection).
		Select("count(*)").
		AndWhere(dbx.HashExp{"owner": ownerId}).
		Row(&total)

	return total, err
}

func isUserSubscribed(dao *daos.Dao, authRecord *models.Record) bool {
	record, _ := dao.FindFirstRecordByData("subscriptions", "user", authRecord.Id)
	return record != nil
}

func main() {
	app := pocketbase.New()

	// check whether a default user has reached the score/file limit
	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)

		if e.Record.Collection().Name != "files" && e.Record.Collection().Name != "scores" {
			return nil
		}

		if admin, _ := e.HttpContext.Get(apis.ContextAdminKey).(*models.Admin); admin != nil {
			return nil
		}

		total, err := findTotalFilesByOwner(app.Dao(), e.Record.Collection(), e.Record.GetString("owner"))
		if err != nil || (!isUserSubscribed(app.Dao(), authRecord) && total >= 2) {
			return apis.NewBadRequestError(fmt.Sprintf("Cannot create more %s!", e.Record.Collection().Name), err)
		}

		return nil
	})

	app.OnRealtimeAfterSubscribeRequest().Add(func(e *core.RealtimeSubscribeEvent) error {
		authRecord, _ := e.Client.Get(apis.ContextAuthRecordKey).(*models.Record)

		if authRecord != nil && len(e.Subscriptions) > 0 {
			fileData, err := app.Dao().FindRecordById("file_data", strings.Split(e.Subscriptions[0], "/")[1])
			if err != nil {
				return err
			}

			existingSubscriptions, err := app.Dao().FindRecordsByExpr("file_data_editors", dbx.HashExp{"file_data": fileData.Id, "user": authRecord.Id})

			if err != nil {
				return err
			}

			if len(existingSubscriptions) > 0 {
				sub := existingSubscriptions[0]
				sub.Set("subscription_id", e.Client.Id())

				if err := app.Dao().SaveRecord(sub); err != nil {
					return err
				}

			} else {
				collection, err := app.Dao().FindCollectionByNameOrId("file_data_editors")
				if err != nil {
					return err
				}

				fileDataEditor := models.NewRecord(collection)
				fileDataEditor.Set("file_data", fileData.Id)
				fileDataEditor.Set("user", authRecord.Id)
				fileDataEditor.Set("subscription_id", e.Client.Id())

				if err := app.Dao().SaveRecord(fileDataEditor); err != nil {
					return err
				}

				editors := append(fileData.GetStringSlice("editors"), fileDataEditor.Id)
				fileData.Set("editors", editors)

				if err := app.Dao().SaveRecord(fileData); err != nil {
					return err
				}
			}
		}

		return nil
	})

	app.OnRealtimeDisconnectRequest().Add(func(e *core.RealtimeDisconnectEvent) error {
		authRecord, _ := e.Client.Get(apis.ContextAuthRecordKey).(*models.Record)

		if authRecord != nil {
			fileDataEditor, err := app.Dao().FindFirstRecordByData("file_data_editors", "subscription_id", e.Client.Id())
			if err != nil {
				return err
			}

			if err != nil {
				return err
			}

			if err := app.Dao().DeleteRecord(fileDataEditor); err != nil {
				return err
			}
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
