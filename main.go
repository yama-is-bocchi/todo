package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nsf/termbox-go"
	"github.com/yama-is-bocchi/todo/app"
	"github.com/yama-is-bocchi/todo/database"
	"github.com/yama-is-bocchi/todo/ui"
)

func main() {
	db, err := database.NewSQLiteDB("db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(db, ui.NewUI(termbox.ColorGreen, termbox.ColorBlack), ui.NewEventHandler())
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
