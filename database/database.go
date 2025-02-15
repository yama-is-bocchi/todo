package database

// データベースインターフェース
type DBInterface interface {
	create(TodoData) error
	read() ([]TodoData, error)
	update(TodoData) error
	delete(TodoData) error
}

// SQLiteを用いたデータベース
type sqliteDatabase struct {
	filePath string
}

// SQLiteを用いたデータベースを提供する
func NewSQLiteDB(filePath string) *sqliteDatabase {
	return &sqliteDatabase{filePath: filePath}
}

// SQLiteにTodoデータ書き込み
func (sqlite *sqliteDatabase) create(targetData TodoData) error {
	return nil
}

// SQLiteからTodoデータ読み込み
func (sqlite *sqliteDatabase) read() ([]TodoData, error) {
	return nil, nil
}

// SQLiteのTodoデータ更新
func (sqlite *sqliteDatabase) update(targetData TodoData) error {
	return nil
}

// SQLiteのTodoデータ削除
func (sqlite *sqliteDatabase) delete(targetData TodoData) error {
	return nil
}
