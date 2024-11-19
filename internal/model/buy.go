package model

import (
	"time"
)

type Buy struct {
	Info      BuyInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type BuyInfo struct {
	Id     int64
	BookId int64
	Amount int64
}
