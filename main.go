package main

import (
	"github.com/astaxie/beego"
	_ "school/models"
	_ "school/routers"
)


func main(){
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}