package service

import (
	"context"

	"github.com/9Neechan/book-store/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (string, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
}


type AuthorService interface {
	Create(ctx context.Context, info *model.AuthorInfo) error
	Get(ctx context.Context, id int) (*model.Author, error)
}

type BookService interface {
	Create(ctx context.Context, info *model.BookInfo) (string, error)
	Get(ctx context.Context, id int) (*model.Book, error)
	UpdatePrice(ctx context.Context, id int)
	UpdateAmount(ctx context.Context, id int)
}

type BuyService interface {
	Create(ctx context.Context, info *model.BuyInfo) error
	Get(ctx context.Context, id int) (*model.Buy, error)
}
