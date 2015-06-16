package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TerminalController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TerminalController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TmpController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TmpController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TermController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TermController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:PageController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:PageController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:CdfController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:CdfController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TemplateController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:TemplateController"],
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

	beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers:HomeController"],
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

}
