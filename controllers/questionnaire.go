package controllers

import (
	"bee/models"
	"fmt"
	"github.com/astaxie/beego"
	orm2 "github.com/astaxie/beego/orm"
)

type QuestionnaireController struct {
	beego.Controller
}

type myStruct struct {
	Name string
}
func (c *QuestionnaireController) Get() {

	c.createUserInfoV1("test", "xxxxxxxxxxxxx")

	var ms myStruct
	ms.Name = "qwsssssssssssssse"
	c.Data["json"] = &ms
	c.ServeJSON()
}

func (c *QuestionnaireController) createUserInfoV1(nickname string, openId string) {
	var user = models.User{Nickname:"xxxx", OpenId:"xxxxxxxxx"}
    orm := orm2.NewOrm()

	_, i2 := orm.Insert(&user)
	if i2 != nil {

		c.Ctx.WriteString("插入失败！" + fmt.Sprint(i2))
	} else {

		c.Ctx.WriteString("插入成功!")
	}

}
