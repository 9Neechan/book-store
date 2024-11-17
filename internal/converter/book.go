package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/9Neechan/book-store/internal/model"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/book"
)

func ToBookFromService(book *model.Book) *desc.Book {
	var updatedAt *timestamppb.Timestamp
	if book.UpdatedAt != nil {
		updatedAt = timestamppb.New(*book.UpdatedAt)
	}

	return &desc.Book{
		Info:      ToBookInfoFromService(book.Info),
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToBookInfoFromService(info model.BookInfo) *desc.BookInfo {
	return &desc.BookInfo{
		Id:       info.Id,
		Title:    info.Title,
		AuthorId: info.AuthorId,
		Genre:    info.Genre,
		Price:    info.Price,
		Amount:   info.Amount,
	}
}

func ToBookInfoFromDesc(info *desc.BookInfo) *model.BookInfo {
	return &model.BookInfo{
		Id:       info.Id,
		Title:    info.Title,
		AuthorId: info.AuthorId,
		Genre:    info.Genre,
		Price:    info.Price,
		Amount:   info.Amount,
	}
}
