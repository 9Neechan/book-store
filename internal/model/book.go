package model

import (
	"time"
)

type Book struct {
	Info      BookInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type BookInfo struct {
	Id       int64
	Title    string
	AuthorId int64
	Genre    string
	Price    int64
	Amount   int64
}
