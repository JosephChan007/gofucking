package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type InfoController struct {
	beego.Controller
}

type User struct {
	Id    int    `form:"-"`
	Name  string `form:"name"`
	Age   int    `form:"age"`
	Email string `form:"email"`
}

func (c *InfoController) Get() {
	c.TplName = "index.html"
}

func (c *InfoController) Post() {
	user := User{}
	if err := c.ParseForm(&user); err != nil {
		panic(err)
	}
	if b, err := json.Marshal(user); err == nil {
		c.Ctx.WriteString(string(b))
	} else {
		panic(err)
	}
}
