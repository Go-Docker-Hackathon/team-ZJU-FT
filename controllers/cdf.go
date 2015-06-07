package controllers

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	"bufio"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	s "strings"
	"time"
)

// Operations about object
type CdfController struct {
	beego.Controller
}

func system(s string) {
	cmd := exec.Command("/bin/sh", "-c", s)
	fmt.Println(s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func Cdfbase() string {
	//base
	strhead, err := ioutil.ReadFile("templatea.txt")
	strlast, err := ioutil.ReadFile("templateb.txt")
	var strtmp string
	var strmid string
	//read line from the history.sh

	//strmid := ioutil.r
	f, err := os.Open("testhis.sh")
	rf := bufio.NewReader(f)
	for {
		byt, _, err := rf.ReadLine()
		if err != nil {
			break
		}
		strtmp = string(byt)
		fmt.Println(strtmp)
		if s.Contains(strtmp, "apt") || s.Contains(strtmp, "git") {
			strmid = strmid + "\n" + "RUN " + strtmp
		} else if s.Contains(strtmp, "vim") || s.Contains(strtmp, "vi") || s.Contains(strtmp, "cat") {
			strcn := s.Split(strtmp, " ")
			if len(strcn) > 0 {
				strmid = strmid + "\n" + "COPY " + strcn[1]
			}

		} else if s.Contains(strtmp, "cd") {
			strcn := s.Split(strtmp, " ")
			if len(strcn) > 0 {
				strmid = strmid + "\n" + "WORKDIR " + s.Split(strtmp, " ")[1]
			}

		} else if s.Contains(strtmp, "ls") || s.Contains(strtmp, "pwd") || s.Contains(strtmp, "exit") {
			continue
		} else {
			strmid = strmid + "\n" + "#" + strtmp
		}
	}

	if err != nil {
		panic(err)
	}

	str := string(strhead) + "\n" + strmid + "\n" + "\n" + string(strlast)

	return str

}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *CdfController) Get() {

	n := 0

	_, err := os.Stat("history.sh")
	if err == nil {
		os.Remove("history.sh")
	}

	fmt.Println("the containerid :", containerid)

	command_commit := `docker commit ` + containerid + " wangzhe:latest" + strconv.Itoa(n)
	cmd := exec.Command("/bin/sh", "-c", command_commit)
	fmt.Println(command_commit)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
	time.Sleep(time.Second)

	//<-done
	fmt.Println("excute")

	//command := `docker run -i wangzhe:latest` + strconv.Itoa(n) + ` /bin/bash -c "cat ~/.bash_history" > history.sh`
	//command = command_commit + comman
	command1 := `docker run -i wangzhe:latest` + strconv.Itoa(n) + ` /bin/bash -c "cat ~/.bash_history" > testhis.sh`
	//system(command)
	system(command1)

	//out put the cmd into the file
	strlast := Cdfbase()
	o.Ctx.WriteString(strlast)

}
