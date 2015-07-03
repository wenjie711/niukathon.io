package models

import (
	"time"
)

type Team struct {
	Id        int64 `xorm:"pk autoincr"`
	Name      string
	Logo      string
	Desc      string
	UpdatedAt time.Time `xorm:"updated_at"`
	CreatedAt time.Time `xorm:"created_at"`
}
