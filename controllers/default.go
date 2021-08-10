package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "liuwei111111@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {
	c.Data["Website"] = "beego.meeeeeeeeeeeeeeeeee"
	c.Data["Email"] = "liuwei111111@gmail.com"
	c.TplName = "index.tpl"
}
