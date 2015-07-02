package controllers

import (
	"github.com/revel/revel"
	"niukathon.io/app/models"
)

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*App).checkUser, revel.BEFORE)

}

func Init() {
	engine = models.Engine
}
