package key_detection

import "github.com/nsf/termbox-go"

// 選択されている項目の行数を返す.
func GetLine(currentLine, lineLength int, event termbox.Event) int {
	switch event.Key {
	case termbox.KeyArrowUp:
		if currentLine > 0 {
			return currentLine - 1
		}
	case termbox.KeyArrowDown:
		if currentLine < lineLength-1 {
			return currentLine + 1
		}
	}
	return currentLine
}

// 終了キーが入力されたか検出する
func IsQuit(event termbox.Event) bool {
	return event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC
}

// Enterキーが入力されたか検出する
func IsEnter(event termbox.Event) bool {
	return event.Key == termbox.KeyEnter
}
