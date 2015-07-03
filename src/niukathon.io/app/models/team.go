package models

import (
	"errors"
	"time"
)

var TeamDao = &Team{}

type Team struct {
	Id        int64  `xorm:"pk autoincr"`
	EventId   int64  `xorm:"notnull unique"`
	UserId    int64  `xorm:"notnull unique"`
	Name      string `xorm:"notnull unique"`
	Logo      string
	Desc      string
	UpdatedAt time.Time `xorm:"updated"`
	CreatedAt time.Time `xorm:"created"`
	DeletedAt time.Time `xorm:"deleted"`
}

var (
	errInternal = errors.New("internal server err")
	errNotFound = errors.New("entity not found")
)

// func (t *Team) CreateTeam(team *Team) (err error) {
// 	affected, err = Engine.Insert(team)
// 	if err != nil {
// 		err = errInternalzz
// 		return
// 	}

// }

func (t *Team) GetTeamList() (teams []Team, err error) {
	teams = make([]Team, 0)
	err = Engine.Find(&teams)
	if err != nil {
		return
	}
	return
}

func (t *Team) GetTeamListByEvent(event *Event) (teams []Team, err error) {
	teams = make([]Team, 0)
	teamCondition := &Team{
		EventId: event.Id,
	}
	err = Engine.Find(&teams, teamCondition)
	if err != nil {
		err = errInternal
		return
	}
	return

}

func (t *Team) GetTeamById(teamCondition *Team) (team Team, err error) {
	has, err := Engine.Get(teamCondition)
	if err != nil {
		err = errInternal
		return
	}
	if !has {
		err = errNotFound
		return
	}

	return
}

func (t *Team) DeleteTeamById(teamCondition *Team) (cnt int64, err error) {
	cnt, err = Engine.Delete(teamCondition)
	if err != nil {
		err = errInternal
		return
	}

	return
}

func (t *Team) UpdateTeam(team *Team) (err error) {
	conditions := map[string]interface{}{
		"Id":      team.Id,
		"EventId": team.EventId,
		"UserId":  team.UserId,
		"Name":    team.Name,
		"Logo":    team.Logo,
		"Desc":    team.Desc,
	}
	_, err = Engine.Table(new(Team)).Id(team.Id).Update(conditions)
	if err != nil {
		err = errInternal
		return
	}
	return
}
