package user

import (
	"context"
	"sync"
	"time"

	"github.com/9Neechan/book-store/internal/model"
	def "github.com/9Neechan/book-store/internal/repository"
	"github.com/9Neechan/book-store/internal/repository/user/converter"
	repoModel "github.com/9Neechan/book-store/internal/repository/user/model"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.User
	m    sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.User),
	}
}

func (r *repository) Create(_ context.Context, userUUID string, info *model.UserInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[userUUID] = &repoModel.User{
		UUID:      userUUID,
		Info:      converter.ToUserInfoFromService(info),
		CreatedAt: time.Now(),
	}

	return nil
}

func (r *repository) Get(_ context.Context, uuid string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	user, ok := r.data[uuid]
	if !ok {
		return nil, nil
	}

	return converter.ToUserFromRepo(user), nil
}
