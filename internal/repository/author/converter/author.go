package converter

import (
	"time"

	"github.com/9Neechan/book-store/internal/model"
	repoModel "github.com/9Neechan/book-store/internal/repository/author/model"
)

func ToAuthorFromRepo(author *repoModel.Author) *model.Author {
	var updatedAt *time.Time
	if author.UpdatedAt.Valid {
		updatedAt = &author.UpdatedAt.Time
	}

	return &model.Author{
		Info:      ToAuthorInfoFromRepo(author.Info),
		CreatedAt: author.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func ToAuthorInfoFromRepo(info repoModel.AuthorInfo) model.AuthorInfo {
	return model.AuthorInfo{
		Id:         info.Id,
		NameAuthor: info.NameAuthor,
	}
}

func ToAuthorInfoFromService(info *model.AuthorInfo) repoModel.AuthorInfo {
	return repoModel.AuthorInfo{
		Id:         info.Id,
		NameAuthor: info.NameAuthor,
	}
}
