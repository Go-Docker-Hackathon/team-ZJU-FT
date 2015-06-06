package main

import (
	_ "github.com/Go-Docker-Hackathon/team-ZJU-FT/docs"
	_ "github.com/Go-Docker-Hackathon/team-ZJU-FT/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
