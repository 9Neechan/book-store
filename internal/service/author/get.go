package author

import (
	"context"
	"log"

	"github.com/9Neechan/book-store/internal/model"
)

func (s *service) Get(ctx context.Context, id int) (*model.Author, error) {
	author, err := s.authorRepository.Get(ctx, id)
	if err != nil {
		log.Printf("ошибка получения автора: %v\n", err)
		return nil, err
	}
	if author == nil {
		log.Printf("автора с id %d не найден\n", id)
		return nil, model.ErrorAuthorNotFound
	}

	return author, nil
}
