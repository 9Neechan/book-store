package app

import (
	"log"

	"github.com/9Neechan/book-store/internal/config"
	"github.com/9Neechan/book-store/internal/repository"
	"github.com/9Neechan/book-store/internal/service"

	"github.com/9Neechan/book-store/internal/api/user"
	userRepository "github.com/9Neechan/book-store/internal/repository/user"
	userService "github.com/9Neechan/book-store/internal/service/user"

	"github.com/9Neechan/book-store/internal/api/author"
	authorRepository "github.com/9Neechan/book-store/internal/repository/author"
	authorService "github.com/9Neechan/book-store/internal/service/author"
)

type ServiceProvider struct {
	grpcConfig config.GRPCConfig
	//pgConfig   config.PGConfig

	userRepository repository.UserRepository
	userService    service.UserService
	userImpl       *user.Implementation

	authorRepository repository.AuthorRepository
	authorService    service.AuthorService
	authorImpl       *author.Implementation
}

func newServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

/*func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}*/

func (s *ServiceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository()
	}

	return s.userRepository
}

func (s *ServiceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}

func (s *ServiceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService())
	}

	return s.userImpl
}

func (s *ServiceProvider) AuthorRepository() repository.AuthorRepository {
	if s.authorRepository == nil {
		s.authorRepository = authorRepository.NewRepository()
	}

	return s.authorRepository
}

func (s *ServiceProvider) AuthorService() service.AuthorService {
	if s.authorService == nil {
		s.authorService = authorService.NewService(
			s.AuthorRepository(),
		)
	}

	return s.authorService
}

func (s *ServiceProvider) AuthorImpl() *author.Implementation {
	if s.authorImpl == nil {
		s.authorImpl = author.NewImplementation(s.AuthorService())
	}

	return s.authorImpl
}
