package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (req *RegisterController) Get() {
	req.TplName = "register.html"
}

func (req *RegisterController) Post() {
	// Register a user
}
