package user

import (
	"github.com/9Neechan/book-store/internal/repository"
	def "github.com/9Neechan/book-store/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}
