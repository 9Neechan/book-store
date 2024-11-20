package author

import (
	repo "github.com/9Neechan/book-store/internal/repository"
	def "github.com/9Neechan/book-store/internal/service"
)

var _ def.AuthorService = (*service)(nil)

type service struct {
	authorRepository repo.AuthorRepository
}

func NewService(
	authorRepository repo.AuthorRepository,
) *service {
	return &service{
		authorRepository: authorRepository,
	}
}
