package book

import (
	service "github.com/9Neechan/book-store/internal/service/book_store"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/book"
)

type Implementation struct {
	desc.UnimplementedBookV1Server
	bookService service.BookService
}

func NewImplementation(bookService service.BookService) *Implementation {
	return &Implementation{
		bookService: bookService,
	}
}
