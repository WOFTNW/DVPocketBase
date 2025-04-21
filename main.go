package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnRecordAfterCreateSuccess("infractions").BindFunc(func(e *core.RecordEvent) error {
		record, err := app.FindRecordById("users", e.Record.Get("user").(string))
		if err != nil {
			app.Logger().Error("Error: Record for User Not Found ", "recordId", record.Id)
			return err
		}
		var s = record.Get("infractions").([]string)
		s = append(s, e.Record.Get("id").(string))
		record.Set("infractions", s)
		err = app.Save(record)
		if err != nil {
			app.Logger().Error("Error Saving Record for User", "recordId", record.Id)
			return err
		}
		return e.Next()
	})
	app.OnRecordAfterCreateSuccess("users_home").BindFunc(func(e *core.RecordEvent) error {
		record, err := app.FindRecordById("users", e.Record.Get("user").(string))
		if err != nil {
			app.Logger().Error("Error: Record for User Not Found ", "recordId", record.Id)
			return err
		}
		var s = record.Get("users_home").([]string)
		s = append(s, e.Record.Get("id").(string))
		record.Set("users_home", s)
		err = app.Save(record)
		if err != nil {
			app.Logger().Error("Error Saving Record for User", "recordId", record.Id)
			return err
		}
		return e.Next()
	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
