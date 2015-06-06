package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {

	h.TplNames = "home.html"
}
