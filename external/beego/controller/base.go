package controller

import (
	"github.com/astaxie/beego"
	"time"
)

type BaseController struct {
	beego.Controller
	code int
	msg  interface{}
	data map[string]interface{}
}

func (p *BaseController) response() {
	result := make(map[string]interface{})
	result["code"] = p.code
	result["msg"] = p.msg
	result["data"] = p.data
	result["server_time"] = time.Now().Format("2006-01-02 15:04:05")

	p.Data["json"] = result
	p.ServeJSON()
}
