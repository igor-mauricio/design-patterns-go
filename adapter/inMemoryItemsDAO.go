package adapter

type InMemoryItemsDAO struct {
	items  map[int64]Item
	lastId int64
}

func NewInMemoryItemsDAO() *InMemoryItemsDAO {
	dao := &InMemoryItemsDAO{
		items:  make(map[int64]Item),
		lastId: -1,
	}
	return dao
}

func (dao *InMemoryItemsDAO) AddItem(name string) (int64, error) {
	dao.lastId++
	dao.items[dao.lastId] = Item{name: name, id: dao.lastId}
	return dao.lastId, nil
}

func (dao *InMemoryItemsDAO) RemoveItem(id int64) error {
	delete(dao.items, id)
	return nil
}

func (dao *InMemoryItemsDAO) ListItems() (map[int64]Item, error) {
	return dao.items, nil
}
