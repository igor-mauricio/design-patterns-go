package adapter

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteItemsDAO struct {
	db *sql.DB
}

func NewSqliteItemsDAO(db *sql.DB) (*SqliteItemsDAO, error) {
	dao := &SqliteItemsDAO{db}
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		return nil, err
	}
	return dao, nil
}

func (dao *SqliteItemsDAO) AddItem(name string) (int64, error) {
	sqlQuery := "INSERT INTO items (name) VALUES (?)"
	result, err := dao.db.Exec(sqlQuery, name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (dao *SqliteItemsDAO) RemoveItem(id int64) error {
	sqlQuery := "DELETE FROM items WHERE ID = ?"
	_, err := dao.db.Exec(sqlQuery, id)
	return err
}

func (dao *SqliteItemsDAO) ListItems() (map[int64]Item, error) {
	sqlQuery := "SELECT * FROM items"
	rows, err := dao.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make(map[int64]Item)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.id, &item.name)
		if err != nil {
			return nil, err
		}
		items[item.id] = item
	}
	return items, nil
}
