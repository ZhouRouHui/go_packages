package router

import (
	"beego_demo/controller"
	"github.com/astaxie/beego"
)

func Test() {

}
func init() {
	beego.Router("/demo", &controller.DemoController{}, "*:Demo")
	beego.Router("/test", &controller.DemoController{}, "*:Test")
}
