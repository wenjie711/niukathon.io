package controllers

const (
	GUEST = iota + 1
	MEMBER
)

var actionPermissions = map[string]int{
	"Site.Index":          GUEST,
	"User.Signin":         GUEST,
	"User.Signout":        GUEST,
	"User.MyProfile":      MEMBER,
	"User.GithubCallBack": GUEST,
	"App.Hello":           MEMBER,
}
