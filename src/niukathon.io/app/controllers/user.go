package controllers

import (
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c *User) SignIn() revel.Result {
	return nil
}

func (c *User) SignOut() revel.Result {
	return nil
}

func (c *User) MyProfile() revel.Result {
	return nil
}
