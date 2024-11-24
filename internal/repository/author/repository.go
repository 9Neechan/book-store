package author

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/9Neechan/book-store/internal/config"
	"github.com/9Neechan/book-store/internal/db"
	"github.com/9Neechan/book-store/internal/model"
	def "github.com/9Neechan/book-store/internal/repository"
	"github.com/9Neechan/book-store/internal/repository/author/converter"
	repoModel "github.com/9Neechan/book-store/internal/repository/author/model"
)

var _ def.AuthorRepository = (*repository)(nil)

type repository struct {
	postgesDb *db.PostgresService
	m         sync.RWMutex
}

func NewRepository() *repository {
	connStr, err := config.GetConnectionString()
	if err != nil {
		log.Fatalf("%v", err)
	}
	pgDB, err := db.NewPostgresDB(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &repository{
		postgesDb: db.NewService(pgDB),
	}
}

func (r *repository) Create(_ context.Context, info *model.AuthorInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	q := fmt.Sprintf("INSERT INTO author (name_author) VALUES %s", info.NameAuthor)

	rows, err := r.postgesDb.Db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *repository) Get(_ context.Context, id int) (*model.Author, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	q := fmt.Sprintf("SELECT * FROM author WHERE id = %d", id)

	rows, err := r.postgesDb.Db.Query(q)
	if err != nil {
		return &model.Author{}, err
	}
	defer rows.Close()

	var author_info *repoModel.AuthorInfo
	var author *repoModel.Author
	if rows.Next() {
		if err := rows.Scan(&author_info.Id, &author_info.NameAuthor, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return &model.Author{}, err
		}
	}

	return converter.ToAuthorFromRepo(author), nil
}
