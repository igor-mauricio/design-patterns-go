package adapter

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryItemsDAO(t *testing.T) {
	dao := NewInMemoryItemsDAO()
	service := NewItemsService(dao)

	items, err := service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 0)

	id1, err := service.AddItem("item1")
	assert.Nil(t, err)
	items, err = service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, id1, int64(0))
	assert.Equal(t, items[id1].name, "item1")

	id2, err := service.AddItem("item2")
	assert.Nil(t, err)
	items, err = service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, items[id2].name, "item2")

	err = service.RemoveItem(id2)
	assert.Nil(t, err)
	items, err = service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 1)
}

func TestSQLiteItemsDAO(t *testing.T) {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec("DROP TABLE items")
	dao, err := NewSqliteItemsDAO(db)
	assert.Nil(t, err)
	service := NewItemsService(dao)

	items, err := service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 0)

	id1, err := service.AddItem("item1")
	assert.Nil(t, err)
	items, err = service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "item1", items[id1].name)

	id2, err := service.AddItem("item2")
	assert.Nil(t, err)
	items, err = service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, "item2", items[id2].name)

	err = service.RemoveItem(id2)
	assert.Nil(t, err)
	items, err = service.ListItems()
	assert.Nil(t, err)
	assert.Len(t, items, 1)
}
