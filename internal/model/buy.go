package model

import (
	"time"
)

type Buy struct {
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type BuyInfo struct {
	Id     int64
	Amount int64
}
