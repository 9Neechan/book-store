package user

import (
	service "github.com/9Neechan/book-store/internal/service/book_store"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/author"
)

type Implementation struct {
	desc.UnimplementedAuthorV1Server
	authorService service.AuthorService
}

func NewImplementation(authorService service.AuthorService) *Implementation {
	return &Implementation{
		authorService: authorService,
	}
}
