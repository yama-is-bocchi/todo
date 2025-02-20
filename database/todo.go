package database

import (
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
// 作成されたTODOデータ
type CreatedData struct {
	Title string
	Desc  string
	Date  date.Date
}
