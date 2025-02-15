package ui

import (
	"github.com/nsf/termbox-go"
)

type UserInterface interface {
	Render(int, []string)
}

type ui struct {
	mainColor termbox.Attribute
	subColor  termbox.Attribute
}

func NewUI(mainColor, subColor termbox.Attribute) *ui {
	return &ui{mainColor: mainColor, subColor: subColor}
}

// 現在の行の情報を反映して引数の行をレンダリングする.
func (ui ui) Render(currentLine int, lines []string) {
	termbox.Clear(ui.mainColor, ui.subColor)
	for i, item := range lines {
		fg := ui.mainColor
		bg := ui.subColor
		if i == currentLine {
			fg, bg = ui.subColor, ui.mainColor
		}
		// setCell
		ui.print(0, i, fg, bg, item)
	}
	termbox.Flush()
}

// 引数の情報で文字列をセルにセットする.
func (ui ui) print(x, y int, fg, bg termbox.Attribute, msg string) {
	for i, r := range msg {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}
