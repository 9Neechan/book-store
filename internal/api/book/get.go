package book

import (
	"context"

	"github.com/9Neechan/book-store/internal/converter"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/book"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequestBook) (*desc.GetResponseBook, error) {
	book, err := i.bookService.Get(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &desc.GetResponseBook{
		Book: converter.ToBookFromService(book),
	}, nil
}
