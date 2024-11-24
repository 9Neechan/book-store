package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/9Neechan/book-store/internal/config"
	"github.com/9Neechan/book-store/internal/db"
	author_desc "github.com/9Neechan/book-store/pkg/book_store/v1/author"
	user_desc "github.com/9Neechan/book-store/pkg/user_v1"
)

type App struct {
	ServiceProvider *ServiceProvider
	grpcServer      *grpc.Server
	postgesDb       *db.PostgresService
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	//err := config.Load(".env")
	err := config.Load("/Users/neechan/dev/go_projects/book-store/.env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.ServiceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	user_desc.RegisterUserV1Server(a.grpcServer, a.ServiceProvider.UserImpl())
	author_desc.RegisterAuthorV1Server(a.grpcServer, a.ServiceProvider.AuthorImpl())

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.ServiceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.ServiceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

/*func (a *App) ConnectDB() {
	pgDB, err := db.NewPostgresDB(a.ServiceProvider.pgConfig.ConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	//defer pgDB.Close()

	a.postgesDb = db.NewService(pgDB)
}*/
