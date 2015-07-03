package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/revel/revel"
	"niukathon.io/app/models"
)

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*App).checkUser, revel.BEFORE)

}

func Init() {
	engine = models.Engine

	githubClientId := revel.Config.StringDefault("github.client_id", "")
	githubClientSecret := revel.Config.StringDefault("github.client_secret", "")
	github = &oauth.Config{
		ClientId:     githubClientId,
		ClientSecret: githubClientSecret,
		AuthURL:      "https://github.com/login/oauth/authorize",
		TokenURL:     "https://github.com/login/oauth/access_token",
		RedirectURL:  "http://niukathon.io/oauth/github",
	}

	revel.INFO.Println(github.AuthCodeURL("haha"))
}
