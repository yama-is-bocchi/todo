package ui

import (
	"github.com/nsf/termbox-go"
)

// ユーザーインターフェース
type UI struct {
	mainColor termbox.Attribute
	subColor  termbox.Attribute
}

// フォントプロパティ
type fontProperty struct {
	x, y            int               // 位置
	fontColor       termbox.Attribute // フォントカラー
	backgroundColor termbox.Attribute // 背景色
	message         string            // メッセージ内容
}

func NewUI(mainColor, subColor termbox.Attribute) *UI {
	return &UI{mainColor: mainColor, subColor: subColor}
}

// 現在の行の情報を反映して引数の行をレンダリングする.
func (ui UI) Render(lines ...fontProperty) {
	termbox.Clear(ui.mainColor, ui.subColor)
	for _, item := range lines {
		// setCell
		ui.print(item)
	}
	termbox.Flush()
}

// 引数の情報で文字列をセルにセットする.
func (ui UI) print(font fontProperty) {
	for i, r := range font.message {
		termbox.SetCell(font.x+i, font.y, r, font.fontColor, font.backgroundColor)
	}
}

// 引数の文字列のfont型を生成する
func (ui UI) GenerateFont(lines ...string) []fontProperty {
	var result []fontProperty
	for y, line := range lines {
		result = append(
			result,
			fontProperty{
				x: 0, y: y,
				fontColor:       ui.mainColor,
				backgroundColor: ui.subColor,
				message:         line,
			})
	}
	return result
}

// 現在位置の行の色を反転する.
func (ui UI) SetCurrentLine(currentLine int, fonts []fontProperty) []fontProperty {
	newFonts := make([]fontProperty, len(fonts))
	copy(newFonts, fonts)
	for i, font := range newFonts {
		if font.y == currentLine {
			newFonts[i] = fontProperty{font.x, font.y, ui.subColor, ui.mainColor, font.message}
			break
		}
	}
	return newFonts
}
