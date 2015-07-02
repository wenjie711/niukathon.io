package controllers

import (
	"github.com/revel/revel"
)

type Site struct {
	App
}

func (c *Site) Index() revel.Result {
	return c.Render()
}

func (c *Site) SignIn() revel.Result {
	return c.Render()
}

func (c *Site) SignOut() revel.Result {
	return c.Render()
}
