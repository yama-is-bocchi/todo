package app

import "github.com/yama-is-bocchi/todo/database"

type mainApplication struct {
	database database.DBInterface[database.TodoData]
}

func New(database database.DBInterface[database.TodoData]) mainApplication {
	return mainApplication{database: database}
}

func (app *mainApplication) Run() {
	println("running...")
}
