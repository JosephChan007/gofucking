package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Stu struct {
	Id      int64   `form:"id"`
	Name    string  `form:"name"`
	Age     int     `form:"age"`
	Bothday string  `form:"bothday"`
	Score   float32 `form:"score"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(hdfs-host4:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(Stu))
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func (this *Stu) Insert() (int64, error) {
	ormer := orm.NewOrm()
	id, err := ormer.Insert(this)
	if err != nil {
		panic(err)
	}
	return id, err
}

func (this *Stu) Update() (int64, error) {
	ormer := orm.NewOrm()
	n, err := ormer.Update(this)
	if err != nil {
		panic(err)
	}

	fmt.Printf("update count: %d\n", n)
	return n, err
}

func (this *Stu) Query() error {
	ormer := orm.NewOrm()
	err := ormer.Read(this)
	if err != nil {
		panic(err)
	}

	n, err1 := json.Marshal(this)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("update count: %d\n", n)

	return err
}
