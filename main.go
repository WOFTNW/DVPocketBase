package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	app.OnRecordAfterCreateSuccess("infractions").BindFunc(func(e *core.RecordEvent) error {
		record, err := app.FindRecordById("users", e.Record.Get("user").(string))
		if err != nil {
			log.Default().Println("Error: Record for User Not Found ", record.Id)
			return err
		}
		var s = record.Get("infractions").([]string)
		s = append(s, e.Record.Get("id").(string))
		record.Set("infractions", s)
		err = app.Save(record)
		if err != nil {
			log.Default().Println("Error Saving Record for User ", record.Id)
			return err
		}
		return e.Next()
	})
}
