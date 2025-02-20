package app

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/yama-is-bocchi/todo/app/internal/key_detection"
	"github.com/yama-is-bocchi/todo/app/internal/screen"
	"github.com/yama-is-bocchi/todo/database"
	"github.com/yama-is-bocchi/todo/ui"
)

type mainApplication struct {
	appDatabase  database.DBInterface
	ui           *ui.UI
	eventHandler *ui.EventHandler
	screenState  screen.ScreenState
}

func New(database database.DBInterface, ui *ui.UI, eventHandler *ui.EventHandler) mainApplication {
	return mainApplication{appDatabase: database, ui: ui, eventHandler: eventHandler, screenState: screen.MENU}
}

func (app *mainApplication) Run() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("failed to initialize termbox:%w", err)
	}
	defer termbox.Close()
	println("running...")
	for {
		switch app.screenState {
		case screen.MENU, screen.LIST:
			state, err := app.printSelectMenu(screen.GetLines(app.screenState)...)
			if err != nil {
				return fmt.Errorf("failed to print select menu:%w", err)
			}
			app.screenState = state
		case screen.CREATE:
			inputs, err := app.printScanMenu(screen.GetLines(app.screenState)...)
			if err != nil {
				return fmt.Errorf("failed to print scan menu:%w", err)
			}
			// 入力画面に対するinputsの処理分け
			fmt.Println(inputs)
			app.screenState = screen.MENU
		}
		if app.screenState == screen.QUIT {
			break
		}
	}
	return nil
}

func (app *mainApplication) printSelectMenu(lines ...string) (screen.ScreenState, error) {
	currentLine := 0
	fonts := app.ui.GenerateFont(lines...)
	for {
		app.ui.Render(app.ui.SetCurrentLine(currentLine, fonts)...)
		event := app.eventHandler.HandleEvent()
		if key_detection.IsQuit(event) {
			break
		}
		if key_detection.IsEnter(event) {
			// 遷移先を取得
			return screen.GetNextState(currentLine, app.screenState), nil
		}
		currentLine = key_detection.GetLine(currentLine, len(lines), event)
	}
	return screen.QUIT, nil
}

func (app *mainApplication) printScanMenu(lines ...string) ([]string, error) {
	var inputs []string
	var runeBuff []rune
	for _, line := range lines {
		app.ui.Render(app.ui.GenerateFont(line)...)
	keyBoardEventLoop:
		for {
			switch ev := app.eventHandler.HandleEvent(); ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeyEnter { // Enterキー
					inputs = append(inputs, string(runeBuff))
					runeBuff = nil
					break keyBoardEventLoop
				}
				if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
					if len(runeBuff) > 0 {
						runeBuff = runeBuff[:len(runeBuff)-1]
					}
				} else if ev.Ch != 0 {
					runeBuff = append(runeBuff, ev.Ch)
				}
				app.ui.Render(app.ui.GenerateFont(line, string(runeBuff))...)
			}
		}
	}
	return inputs, nil
}
