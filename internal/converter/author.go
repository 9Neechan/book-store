package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/9Neechan/book-store/internal/model"
	desc "github.com/9Neechan/book-store/pkg/book_store/v1/author"
)

func ToAuthorFromService(author *model.Author) *desc.Author {
	var updatedAt *timestamppb.Timestamp
	if author.UpdatedAt != nil {
		updatedAt = timestamppb.New(*author.UpdatedAt)
	}

	return &desc.Author{
		Info:      ToAuthorInfoFromService(author.Info),
		CreatedAt: timestamppb.New(author.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToAuthorInfoFromService(info model.AuthorInfo) *desc.AuthorInfo {
	return &desc.AuthorInfo{
		Id:         info.Id,
		NameAuthor: info.NameAuthor,
	}
}

func ToAuthorInfoFromDesc(info *desc.AuthorInfo) *model.AuthorInfo {
	return &model.AuthorInfo{
		Id:         info.Id,
		NameAuthor: info.NameAuthor,
	}
}
