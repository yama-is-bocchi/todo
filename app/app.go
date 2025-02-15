package app

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/yama-is-bocchi/todo/app/internal/key_detection"
	"github.com/yama-is-bocchi/todo/database"
	"github.com/yama-is-bocchi/todo/ui"
)

type mainApplication struct {
	appDatabase  database.DBInterface
	ui           ui.UserInterface
	eventHandler ui.EventHandlerInterface
}

func New(database database.DBInterface, ui ui.UserInterface, eventHandler ui.EventHandlerInterface) mainApplication {
	return mainApplication{appDatabase: database, ui: ui, eventHandler: eventHandler}
}

func (app *mainApplication) Run() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("failed to initialize termbox:%w", err)
	}
	println("running...")
	currentLine := 0
	lines := []string{"テスト項目1", "テスト項目2", "テスト項目3"}
	for {
		app.ui.Render(currentLine, lines)
		event := app.eventHandler.HandleEvent()
		if key_detection.IsQuit(event) {
			termbox.Close()
			break
		}
		if key_detection.IsEnter(event) {
			// 次の画面に遷移or何か処理
			fmt.Println("Entered!")
		}
		currentLine = key_detection.GetLine(currentLine, len(lines), event)
	}
	return nil
}
