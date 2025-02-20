package screen

type ScreenState int

const (
	MENU ScreenState = iota
	CREATE
	LIST
	EDIT
	DELETE
	QUIT
)

func GetNextState(currentLine int, currentState ScreenState) ScreenState {
	switch currentState {
	case MENU:
		switch currentLine {
		case 0:
			return CREATE
		case 1:
			return LIST
		case 2:
			return EDIT
		case 3:
			return DELETE
		}
	}
	return MENU
}

func GetLines(state ScreenState) []string {
	switch state {
	case MENU:
		return []string{"1.登録", "2.リスト", "3.編集", "4.削除"}
	case CREATE:
		return []string{"タイトル", "説明", "日付"}
	}
	return nil
}
