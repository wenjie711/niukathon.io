package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/revel/revel"
)

type User struct {
	App
}

func (c *User) GithubCallBack(code string) revel.Result {
	t := &oauth.Transport{Config: github}
	tok, err := t.Exchange(code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect("/")
	}
	c.Session["access_token"] = t.AccessToken

	revel.INFO.Println("tok", tok)

	return c.Redirect("/")
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
