package controllers

import (
	"github.com/astaxie/beego/orm"
	"school/models"
)


type MainController struct {
	Base
}

func (c *MainController) Get() {
	set := models.Dataset{}
	set.Id = 1
	o := orm.NewOrm()
	o.Read(&set)
	c.Data["onlytitle"] = set.Onlytitle
	c.TplName = "Index/index.html"
}
