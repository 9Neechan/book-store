package buy

import (
	"context"

	"github.com/9Neechan/book-store/internal/converter"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/buy"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	buy, err := i.buyService.Get(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Buy: converter.ToBuyFromService(buy),
	}, nil
}
