package models

import (
	"encoding/json"
	"github.com/revel/revel"
	"net/http"
	"net/url"
	"time"
)

type User struct {
	Id        int64     `xorm:"pk autoincr"`
	Username  string    `json:"login"`
	Nickname  string    `json:"name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar_url"`
	UpdatedAt time.Time `xorm:"updated_at"`
	CreatedAt time.Time `xorm:"created_at"`
}

func GithubAuthenticated(access_token string) (*User, error) {
	resp, _ := http.Get("https://api.github.com/user?access_token=" +
		url.QueryEscape(access_token))
	defer resp.Body.Close()
	user := new(User)
	err := json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		revel.ERROR.Println(err)
		return user, err
	}
	revel.INFO.Println(user)
	has, _ := Engine.Where("username=?", user.Username).Get(user)
	if !has {
		revel.INFO.Println(user.Username, " not found")
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt
		user.Id, err = Engine.InsertOne(user)
	}
	return user, err
}

func (u User) Login() bool {
	return u.Username != "" && u.Id > 0
}
