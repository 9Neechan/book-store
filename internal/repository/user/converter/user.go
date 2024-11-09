package converter

import (
	"time"

	"github.com/9Neechan/book-store/internal/model"
	repoModel "github.com/9Neechan/book-store/internal/repository/user/model"
)

func ToUserFromRepo(user *repoModel.User) *model.User {
	var updatedAt *time.Time
	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}

	return &model.User{
		UUID:      user.UUID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(info repoModel.UserInfo) model.UserInfo {
	return model.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Age:       info.Age,
	}
}

func ToUserInfoFromService(info *model.UserInfo) repoModel.UserInfo {
	return repoModel.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Age:       info.Age,
	}
}
