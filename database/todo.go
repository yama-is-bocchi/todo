package database

import (
	"fmt"

	"github.com/thamaji/date"
)

// TODOデータ構造
type TodoData struct {
	serialNumber    int
	title           string
	description     string
	limitDate       date.Date
	createdDate     date.Date
	lastUpdatedDate date.Date
}

func (todo TodoData) String() string {
	return fmt.Sprintf(`
*** %d ***

*** タイトル ***
%s

*** 説明 ***
%s

*** 期限 ***
%s

*** 作成日 ***
%s

*** 最終更新日 ***
%s

`, todo.serialNumber, todo.title, todo.description, todo.limitDate, todo.createdDate.String(), todo.lastUpdatedDate.String())
}

// 作成されたTODOデータ
type CreatedData struct {
	Title string
	Desc  string
	Date  date.Date
}
