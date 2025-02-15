package fake

import (
	"errors"

	"github.com/timut2/avito_test_task/internal/repository"
)

var ErrItemNotFound = errors.New("item not found")

type fakeItemRepo struct {
	items map[uint64]ItemInfo
}

type ItemInfo struct {
	Name  string
	Price int
}

func NewFakeItemRepo() *fakeItemRepo {
	return &fakeItemRepo{
		items: map[uint64]ItemInfo{
			1:  {Name: "t-shirt", Price: 80},
			2:  {Name: "cup", Price: 20},
			3:  {Name: "book", Price: 50},
			4:  {Name: "pen", Price: 10},
			5:  {Name: "powerbank", Price: 200},
			6:  {Name: "hoody", Price: 300},
			7:  {Name: "umbrella", Price: 200},
			8:  {Name: "socks", Price: 10},
			9:  {Name: "wallet", Price: 50},
			10: {Name: "pink-hoody", Price: 500},
		},
	}
}

func (r *fakeItemRepo) GetCostByID(itemID int) (int, error) {
	item, exists := r.items[uint64(itemID)]
	if !exists {
		return 0, repository.ErrItemNotFound
	}
	return item.Price, nil
}
