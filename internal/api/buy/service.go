package buy

import (
	service "github.com/9Neechan/book-store/internal/service/book_store"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/buy"
)

type Implementation struct {
	desc.UnimplementedBuyV1Server
	buyService service.BuyService
}

func NewImplementation(buyService service.BuyService) *Implementation {
	return &Implementation{
		buyService: buyService,
	}
}
