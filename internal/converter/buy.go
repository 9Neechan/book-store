package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/9Neechan/book-store/internal/model"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/buy"
)

func ToBuyFromService(buy *model.Buy) *desc.Buy {
	var updatedAt *timestamppb.Timestamp
	if buy.UpdatedAt != nil {
		updatedAt = timestamppb.New(*buy.UpdatedAt)
	}

	return &desc.Buy{
		Info:      ToBuyInfoFromService(buy.Info),
		CreatedAt: timestamppb.New(buy.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToBuyInfoFromService(info model.BuyInfo) *desc.BuyInfo {
	return &desc.BuyInfo{
		Id:     info.Id,
		BookId: info.BookId,
		Amount: info.Amount,
	}
}

func ToBuyInfoFromDesc(info *desc.BuyInfo) *model.BuyInfo {
	return &model.BuyInfo{
		Id:     info.Id,
		BookId: info.BookId,
		Amount: info.Amount,
	}
}
