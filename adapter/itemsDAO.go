package adapter

type Item struct {
	name string
	id   int64
}

type ItemsDAO interface {
	AddItem(name string) (int64, error)
	RemoveItem(id int64) error
	ListItems() (map[int64]Item, error)
}
