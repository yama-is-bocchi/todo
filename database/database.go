package database

// データベースインターフェース
type DBInterface[T any] interface {
	write([]T) error
	read() ([]T, error)
}

// SQLiteを用いたデータベース
type sqliteDatabase struct {
	filePath string
}

// SQLiteを用いたデータベースを提供する
func NewSQLiteDB(filePath string) *sqliteDatabase {
	return &sqliteDatabase{filePath: filePath}
}

// SQLiteからTodoデータ読み込み
func (sqlite *sqliteDatabase) read() ([]TodoData, error) {
	return nil, nil
}

// SQLiteにTodoデータ書き込み
func (sqlite *sqliteDatabase) write([]TodoData) error {
	return nil
}
