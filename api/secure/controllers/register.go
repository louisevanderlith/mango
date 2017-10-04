package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (req *RegisterController) Get() {
	// Returns the Registration Page
}

func (req *RegisterController) Post() {
	// Register a user
}
