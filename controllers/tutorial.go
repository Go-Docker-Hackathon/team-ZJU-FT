package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

type TutorialController struct {
	beego.Controller
}

func init() {
	session, err := mgo.Dial("121.40.236.126")
	if err != nil {
		panic(err)
	}
	//	defer session.Close() // Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Strong, true)
	c = session.DB("toturial").C("cmd")
	lens = len(names)
	fmt.Println("connected")
}

var names = []string{"FROM", "MAINTAINER", "USER", "LABEL", "ENV", "WORKDIR", "RUN", "COPY", "ADD", "VOLUME", "ONBUILD", "EXPOSE", "ENTRYPOINT", "CMD"}
var index = 0
var cmd = Command{}
var c *mgo.Collection
var lens int

// @Title render main page
// @Description : start the websocket connection
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (this *TutorialController) Get() {
	tutname := this.Input().Get("tutname")
	for i, name := range names {
		if strings.ToUpper(tutname) == name {
			index = i
			break
		}
	}
	log.Printf("tutname is %v , index is %v", tutname, index)
	err := c.Find(bson.M{"type": names[index]}).One(&cmd)
	response := make(map[string]string)
	if err != nil {
		log.Fatal(err)
		response["code"] = "404"
		this.Data["json"] = response
		this.ServeJson()
		return
	}
	log.Println(cmd)
	response["title"] = cmd.Type
	response["content"] = cmd.Example
	if index > 0 && index < lens-1 {
		response["prev"] = names[index-1]
		response["next"] = names[index+1]
	} else if index == 0 {
		response["prev"] = names[lens-1]
		response["next"] = names[index+1]
	} else {
		response["prev"] = names[index-1]
		response["next"] = names[0]
	}
	response["code"] = "200"
	this.Data["json"] = response
	this.ServeJson()
}
