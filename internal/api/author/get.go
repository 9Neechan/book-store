package author

import (
	"context"

	"github.com/9Neechan/book-store/internal/converter"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/author"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequestAuthor) (*desc.GetResponseAuthor, error) {
	author, err := i.authorService.Get(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &desc.GetResponseAuthor{
		Author: converter.ToAuthorFromService(author),
	}, nil
}
