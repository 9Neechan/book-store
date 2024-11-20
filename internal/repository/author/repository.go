package author

import (
	"context"
	"sync"
	"time"

	"github.com/9Neechan/book-store/internal/model"
	def "github.com/9Neechan/book-store/internal/repository"
	"github.com/9Neechan/book-store/internal/repository/author/converter"
	repoModel "github.com/9Neechan/book-store/internal/repository/author/model"
)

var _ def.AuthorRepository = (*repository)(nil)

type repository struct {
	data map[int]*repoModel.Author
	m    sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[int]*repoModel.Author),
	}
}

func (r *repository) Create(_ context.Context, info *model.AuthorInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	id := 0

	r.data[id] = &repoModel.Author{
		Info:      converter.ToAuthorInfoFromService(info),
		CreatedAt: time.Now(),
	}

	return nil
}

func (r *repository) Get(_ context.Context, id int) (*model.Author, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	author, ok := r.data[id]
	if !ok {
		return nil, nil
	}

	return converter.ToAuthorFromRepo(author), nil
}
