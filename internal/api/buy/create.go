package buy

import (
	"context"

	"github.com/9Neechan/book-store/internal/converter"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/buy"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequestBuy) (*desc.CreateResponseBuy, error) {
	err := i.buyService.Create(ctx, converter.ToBuyInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponseBuy{
	}, nil
}

