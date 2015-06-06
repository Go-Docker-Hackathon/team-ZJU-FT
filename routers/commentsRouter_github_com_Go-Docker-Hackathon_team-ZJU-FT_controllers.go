package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:BuildController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:BuildController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:BuildController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:PageController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:PageController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TestController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TutorialController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TutorialController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
