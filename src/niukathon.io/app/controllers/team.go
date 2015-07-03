package controllers

import (
	"github.com/revel/revel"
	"niukathon.io/app/models"
)

type Team struct {
	*revel.Controller
}

func (c *Team) GetTeam(id int64) revel.Result {
	team := &models.Team{
		Id: id,
	}
	models.TeamDao.GetTeamById(team)
	return c.Render(team)
}
