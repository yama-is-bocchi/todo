package database

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
func NewSQLiteDB(filePath string) *sqliteDatabase {
	return &sqliteDatabase{filePath: filePath}
}

// SQLiteにTodoデータ書き込み
func (sqlite *sqliteDatabase) Create(targetData CreatedData) error {
	return nil
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
