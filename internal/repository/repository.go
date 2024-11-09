package repository

import (
	"context"

	"github.com/9Neechan/book-store/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	Get(ctx context.Context, uuid string) (*model.User, error)
}
