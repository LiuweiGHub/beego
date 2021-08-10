package routers

import (
	"bee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Get("/",func(ctx *context.Context){
	//	ctx.Output.Body([]byte("hello world"))
	//})
    beego.Router("/", &controllers.MainController{}, "get:Test")
    beego.Router("/test", &controllers.QuestionnaireController{})
}
