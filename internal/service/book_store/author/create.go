package author

import (
	"context"
	"log"

	"github.com/9Neechan/book-store/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.AuthorInfo) (error) {
	err := s.authorRepository.Create(ctx, info)
	if err != nil {
		log.Printf("ошибка создания пользователя: %v\n", err)
		return err
	}

	return nil
}