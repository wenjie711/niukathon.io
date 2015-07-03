package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"niukathon.io/app/models"
)

var (
	engine *xorm.Engine
)

var github *oauth.Config

type App struct {
	*revel.Controller
}

func (c *App) checkUser() revel.Result {
	user := c.user()
	if user == nil {
		user = new(models.User)
	}
	revel.INFO.Println(user)
	c.RenderArgs["user"] = user

	c.RenderArgs["github"] = github.AuthCodeURL("haha")
	if needCheck, ok := actionPermissions[c.Action]; ok && needCheck > GUEST && user == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/signin?continue=%s", c.Request.Request.URL.String())
	}

	return nil
}

func (c *App) user() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}

	if username, ok := c.Session["user"]; ok {
		return c.getUserFromDB(username)
	}

	if access_token, ok := c.Session["access_token"]; ok {
		user, err := models.GithubAuthenticated(access_token)
		if err != nil {
			return nil
		} else {
			c.Session["user"] = user.Username
			return user
		}
	}

	return nil
}

func (c *App) getUserFromDB(username string) *models.User {
	var user models.User
	has, _ := engine.Where("username = ?", username).Get(&user)
	if !has {
		return nil
	}

	return &user
}
