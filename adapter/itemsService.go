package adapter

type ItemsService struct {
	dao ItemsDAO
}

func NewItemsService(dao ItemsDAO) *ItemsService {
	service := &ItemsService{dao}
	return service
}

func (service *ItemsService) AddItem(name string) (int64, error) {
	return service.dao.AddItem(name)
}

func (service *ItemsService) RemoveItem(id int64) error {
	return service.dao.RemoveItem(id)
}

func (service *ItemsService) ListItems() (map[int64]Item, error) {
	return service.dao.ListItems()
}
