package models

import (
	"time"
)

type Post struct {
	Id        int64
	Uid       int64 `xorm:"user_id"`
	Content   string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}
