package main

import (
	"fmt"
	"log"

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

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
