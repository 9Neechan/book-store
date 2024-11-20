package author

import (
	"context"

	"github.com/9Neechan/book-store/internal/converter"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/author"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequestAuthor) (*desc.CreateResponseAuthor, error) {
	err := i.authorService.Create(ctx, converter.ToAuthorInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponseAuthor{
	}, nil
}
