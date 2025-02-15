package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/yama-is-bocchi/todo/app"
	"github.com/yama-is-bocchi/todo/database"
)

func main() {
	app := app.New(database.NewSQLiteDB("db.sqlite"))
	app.Run()
}
