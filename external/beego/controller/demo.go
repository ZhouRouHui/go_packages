package controller

type DemoController struct {
	BaseController
}

func (p *DemoController) Demo() {
	defer p.response()

	p.code = 1001
	p.msg = "ahakid"
	p.data = map[string]interface{}{
		"a": "a",
		"b": "b",
	}
}

func (p *DemoController) Test() {
	p.Data["json"] = map[string]interface{}{
		"a": "a",
		"b": "b",
	}
	p.ServeJSON()
}
