package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	conn *sql.DB
}

// NewPostgresDB - создание подключения к PostgreSQL
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresDB{conn: db}, nil
}

func (p *PostgresDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.conn.Query(query, args...)
}

func (p *PostgresDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return p.conn.Exec(query, args...)
}

func (p *PostgresDB) Close() error {
	return p.conn.Close()
}

