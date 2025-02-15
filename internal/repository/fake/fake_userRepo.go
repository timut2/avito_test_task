package fake

import (
	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository"
)

type fakeUserRepo struct {
	users        map[int64]*models.User
	usernameToID map[string]int
	currID       int64
}

func NewFakeUserRepo() *fakeUserRepo {
	return &fakeUserRepo{
		users:        make(map[int64]*models.User),
		usernameToID: make(map[string]int),
		currID:       1,
	}
}

func (r *fakeUserRepo) Create(user *models.User) (int, error) {
	user.Id = int(r.currID)
	r.users[r.currID] = user
	r.usernameToID[user.Username] = int(r.currID)
	r.currID++
	return int(r.currID) - 1, nil
}

func (r *fakeUserRepo) FindByUserName(username string) (*models.User, error) {
	if id, ok := r.usernameToID[username]; ok {
		if user, ok := r.users[int64(id)]; ok {
			return user, nil
		}
	}
	return nil, repository.ErrUserNotFound
}

func (r *fakeUserRepo) BalanceById(userId int) (int, error) {
	if user, ok := r.users[int64(userId)]; ok {
		return user.Coins, nil
	}
	return 0, repository.ErrUserNotFound
}
