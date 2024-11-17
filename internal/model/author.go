package model

import (
	"time"
)

type Author struct {
	Info      AuthorInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type AuthorInfo struct {
	Id         int64
	NameAuthor string
}
