package db

import "database/sql"

// Database - общий интерфейс для баз данных
type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

/*func New(db Database) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db Database
}*/