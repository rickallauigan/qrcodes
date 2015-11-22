package controllers

import (
	"github.com/astaxie/beegae"
)

type MainController struct {
	beegae.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "techaguru.com"
	this.Data["Email"] = "rick@techaguru.com"
	this.TplNames = "main.html"
}
