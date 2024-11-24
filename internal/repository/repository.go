package repository

import (
	"context"

	"github.com/9Neechan/book-store/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	Get(ctx context.Context, uuid string) (*model.User, error)
} 

type AuthorRepository interface {
	Create(ctx context.Context, info *model.AuthorInfo) error
	Get(ctx context.Context, authorID int) (*model.Author, error)
} 

type BookRepository interface {
	Create(ctx context.Context, info *model.BookInfo) (string, error)
	Get(ctx context.Context, id int) (*model.Book, error)
	UpdatePrice(ctx context.Context, id int)
	UpdateAmount(ctx context.Context, id int)
}

type BuyRepository interface {
	Create(ctx context.Context, info *model.BuyInfo) (error)
	Get(ctx context.Context, id int) (*model.Buy, error)
}