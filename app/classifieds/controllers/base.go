package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	HasScript  bool
	ScriptName string
}

func (this *BaseController) Setup(name string) {
	this.Layout = "master.html"
	this.TplName = "content/" + name + ".html"

	// By default we want to include scripts
	// Set this to false in your controller, when scripts aren't needed
	this.Data["HasScript"] = true
	this.Data["ScriptName"] = name + ".entry.js"
}
