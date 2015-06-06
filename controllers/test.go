package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
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
func (u *TestController) Get() {
	fmt.Println("Yes")
}
