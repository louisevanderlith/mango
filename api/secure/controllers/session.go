package controllers

import (
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

// @Title Get User Info
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func Get() {
	// return the user's roles
	
}