package database

import (
	"database/sql"
	"fmt"
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
	// データベースに接続 (なければ作成)
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
	return nil, nil
}

// SQLiteのTodoデータ更新
func (sqlite *sqliteDatabase) Update(targetData TodoData) error {
	return nil
}

// SQLiteのTodoデータ削除
func (sqlite *sqliteDatabase) Delete(targetData TodoData) error {
	return nil
}
