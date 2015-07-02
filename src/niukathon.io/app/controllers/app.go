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

var CONFIG = &oauth.Config{
	ClientId:     "95341411595",
	ClientSecret: "8eff1b488da7fe3426f9ecaf8de1ba54",
	AuthURL:      "http://account.qiniu.com/oauth/authorize",
	TokenURL:     "http://account.qiniu.com/oauth/token",
	RedirectURL:  "http://app1.qiniu.com:9011/Application/Auth",
}

type App struct {
	*revel.Controller
}

func (c *App) checkUser() revel.Result {
	user := c.user()
	if user != nil {
		c.RenderArgs["user"] = user
	}
	revel.INFO.Println("guest", GUEST, "member", MEMBER, user)
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
		return c.getUser(username)
	}

	return nil
}

func (c *App) getUser(username string) *models.User {
	var user models.User
	has, _ := engine.Where("username = ?", username).Get(&user)
	if !has {
		return nil
	}

	return &user
}

func (c *App) Index() revel.Result {
	return c.Render()
}

func (c *App) Hello() revel.Result {
	return nil
}
