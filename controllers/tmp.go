package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

type TmpController struct {
	beego.Controller
}

var dockerfile_names = []string{"redis", "mongodb", "apache", "mysql"}
var dockerfile Dockerfile

// @Title render main page
// @Description : start the websocket connection
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (this *TmpController) Get() {
	name := this.Input().Get("name")
	for i, name := range dockerfile_names {
		if strings.ToLower(name) == name {
			index = i
			break
		}
	}
	log.Printf("dockerfile name is %v , index is %v", name, index)
	err := c.Find(bson.M{"type": dockerfile_names[index]}).One(&dockerfile)
	response := make(map[string]string)
	if err != nil {
		log.Fatal(err)
		response["dockerfile"] = "404"
		this.Data["json"] = response
		this.ServeJson()
		return
	}
	log.Println(dockerfile)
	response["dockerfile"] = dockerfile.Template
	this.Data["json"] = response
	this.ServeJson()
}
