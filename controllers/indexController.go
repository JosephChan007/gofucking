package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	fmt.Println("[*] hello hello hello world!")
	c.Ctx.WriteString("[*]hello hello world world!!")
}

func (c *IndexController) Get() {
	fmt.Println("[Get] hello hello hello world!")
	c.Ctx.WriteString("[Get]hello hello world world!!")
}

func (c *IndexController) Post() {
	fmt.Println("[Post] hello hello hello world!")
	c.Ctx.WriteString("[Post]hello hello world world!!")
}

func (c *IndexController) Book() {
	id := c.GetString("id")
	name := c.Input().Get("name")
	fmt.Printf("id:%s, name:%s\n", id, name)
	c.Ctx.ResponseWriter.Write([]byte(id + "\n"))
	c.Ctx.WriteString(name)
}
