package ui

import (
	"github.com/nsf/termbox-go"
)

type UserInterface interface {
	Render(int, []string)
}

// ユーザーインターフェース
type ui struct {
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

func NewUI(mainColor, subColor termbox.Attribute) *ui {
	return &ui{mainColor: mainColor, subColor: subColor}
}

// 現在の行の情報を反映して引数の行をレンダリングする.
func (ui ui) Render(currentLine int, lines []string) {
	termbox.Clear(ui.mainColor, ui.subColor)
	fonts := ui.setCurrentLine(currentLine, ui.generateFont(lines...))
	for _, item := range fonts {
		// setCell
		ui.print(item)
	}
	termbox.Flush()
}

// 引数の情報で文字列をセルにセットする.
func (ui ui) print(font fontProperty) {
	for i, r := range font.message {
		termbox.SetCell(font.x+i, font.y, r, font.fontColor, font.backgroundColor)
	}
}

// 引数の文字列のfont型を生成する
func (ui ui) generateFont(lines ...string) []fontProperty {
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
func (ui ui) setCurrentLine(currentLine int, fonts []fontProperty) []fontProperty {
	for i, font := range fonts {
		if font.y == currentLine {
			fonts[i] = fontProperty{font.x, font.y, font.backgroundColor, font.fontColor, font.message}
			break
		}
	}
	return fonts
}
