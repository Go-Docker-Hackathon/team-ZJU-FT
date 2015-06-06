package controllers

import (
	"github.com/Go-Docker-Hackathon/team-ZJU-FT/builder/parser"
	"github.com/astaxie/beego"
	"strings"
)

type TestController struct {
	beego.Controller
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (this *TestController) Get() {
	dockerfile := this.Input().Get("dockerfile")
	_, err := parser.Parse(strings.NewReader(dockerfile))
	response := make(map[string]string)
	if err != nil {
		response["code"] = "200"
		response["result"] = err.Error()
	} else {
		response["code"] = "404"
		response["result"] = "Congratulations!"
	}

	this.Data["json"] = response
	this.ServeJson()
}
