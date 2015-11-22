package routers

import (
	"github.com/astaxie/beegae"
	"techaguru-orangepineapple/controllers"
)

func init() {
	beegae.Router("/", &controllers.MainController{})
}
