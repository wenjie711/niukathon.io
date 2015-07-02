package models

import (
	"time"
)

type User struct {
	Id        int64 `xorm:"pk autoincr"`
	Username  string
	Nickname  string
	Email     string
	Avatar    string
	UpdatedAt time.Time `xorm:"updated_at"`
	CreatedAt time.Time `xorm:"created_at"`
}
