package model

import (
	"database/sql"
	"time"
)

type Author struct {
	Info      AuthorInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type AuthorInfo struct {
	Id         int64
	NameAuthor string
}
