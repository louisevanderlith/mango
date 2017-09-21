package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Setup("home")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
}
