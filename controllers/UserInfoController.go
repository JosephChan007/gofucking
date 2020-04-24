package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gofucking/models"
)

type UserInfoController struct {
	beego.Controller
}

func (c *UserInfoController) Get() {
	c.TplName = "stu.html"
}

func (c *UserInfoController) Insert() {
	student := models.Stu{}
	if err := c.ParseForm(&student); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", student)

	id, err := student.Insert()
	if err != nil {
		panic(err)
	}

	student.Id = id
	s, err := json.Marshal(&student)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
	c.Ctx.WriteString(string(s))
}

func (c *UserInfoController) Update() {
	student := models.Stu{}
	if err := c.ParseForm(&student); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", student)

	_, err := student.Update()
	if err != nil {
		panic(err)
	}

	s, err := json.Marshal(&student)
	if err != nil {
		panic(err)
	}

	c.Ctx.WriteString(string(s))
}

func (c *UserInfoController) Query() {
	student := models.Stu{Id: 3}
	err := student.Query()
	if err != nil {
		panic(err)
	}

	s, err := json.Marshal(&student)
	if err != nil {
		panic(err)
	}

	c.Ctx.WriteString(string(s))
}
