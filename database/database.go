package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/thamaji/date"
)

// データベースインターフェース
type DBInterface interface {
	Create(CreatedData) error
	Read() ([]TodoData, error)
	Update(TodoData) error
	Delete(TodoData) error
}

// SQLiteを用いたデータベース
type sqliteDatabase struct {
	filePath string
}

// SQLiteを用いたデータベースを提供する
func NewSQLiteDB(filePath string) (*sqliteDatabase, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite3:%w", err)
	}
	defer db.Close()
	// テーブル作成
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
    serialNumber    INTEGER PRIMARY KEY AUTOINCREMENT,
    title           TEXT NOT NULL,
    description     TEXT NOT NULL,
    limitDate       TEXT NOT NULL,
    createdDate     TEXT NOT NULL,
    lastUpdatedDate TEXT NOT NULL
);
	`
	_, err = db.Exec(createTableSQL)
	return &sqliteDatabase{filePath: filePath}, err
}

// SQLiteにTodoデータ書き込み
func (sqlite *sqliteDatabase) Create(targetData CreatedData) error {
	db, err := sql.Open("sqlite3", sqlite.filePath)
	if err != nil {
		return fmt.Errorf("failed to open sqlite3:%w", err)
	}
	defer db.Close()
	query := `
	INSERT INTO todos (title, description, limitDate, createdDate, lastUpdatedDate)
	VALUES (?, ?, ?, datetime('now'), datetime('now'));
	`
	_, err = db.Exec(query, targetData.Title, targetData.Desc, targetData.Date.String())
	return err
}

// SQLiteからTodoデータ読み込み
func (sqlite *sqliteDatabase) Read() ([]TodoData, error) {
	db, err := sql.Open("sqlite3", sqlite.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite3: %w", err)
	}
	defer db.Close()

	query := `SELECT serialNumber, title, description, limitDate, createdDate, lastUpdatedDate FROM todos`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []TodoData

	for rows.Next() {
		var todo TodoData
		var limitDateStr, createdDateStr, lastUpdatedDateStr string

		err := rows.Scan(
			&todo.serialNumber, &todo.title, &todo.description,
			&limitDateStr, &createdDateStr, &lastUpdatedDateStr,
		)
		if err != nil {
			return nil, err
		}

		if limitDateStr != "" {
			t, err := time.Parse("2006-01-02", limitDateStr)
			if err != nil {
				log.Println("warning: failed to parse limitDate:", err)
			} else {
				todo.limitDate = date.FromTime(t)
			}
		}

		if createdDateStr != "" {
			t, err := time.Parse("2006-01-02 15:04:05", createdDateStr)
			if err != nil {
				log.Println("warning: failed to parse createdDate:", err)
			} else {
				todo.createdDate = date.FromTime(t)
			}
		}

		if lastUpdatedDateStr != "" {
			t, err := time.Parse("2006-01-02 15:04:05", lastUpdatedDateStr)
			if err != nil {
				log.Println("warning: failed to parse lastUpdatedDate:", err)
			} else {
				todo.lastUpdatedDate = date.FromTime(t)
			}
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

// SQLiteのTodoデータ更新
func (sqlite *sqliteDatabase) Update(targetData TodoData) error {
	return nil
}

// SQLiteのTodoデータ削除
func (sqlite *sqliteDatabase) Delete(targetData TodoData) error {
	return nil
}
