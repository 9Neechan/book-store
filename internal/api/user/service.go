package user

import (
	"github.com/9Neechan/book-store/internal/service"
	desc "github.com/9Neechan/book-store/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
