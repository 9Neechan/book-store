package book

import (
	"context"
	//"strconv"

	"github.com/9Neechan/book-store/internal/converter"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/book"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequestBook) (*desc.CreateResponseBook, error) {
	_, err := i.bookService.Create(ctx, converter.ToBookInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponseBook{
		
	}, nil
}
