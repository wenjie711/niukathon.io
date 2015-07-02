package models

import (
	"time"
)

type Event struct {
	Id        int64 `xorm:"pk autoincr"`
	Name      string
	Desc      string
	Logo      string
	StartTime time.Time `xorm:"start_time"`
	EndTime   time.Time `xorm:"end_time"`
	Status    int
	UpdatedAt time.Time `xorm:"updated_at"`
	CreatedAt time.Time `xorm:"created_at"`
}
