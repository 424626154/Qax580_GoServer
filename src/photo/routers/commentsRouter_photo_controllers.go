package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["photo/controllers:PController"] = append(beego.GlobalControllerRouter["photo/controllers:PController"],
		beego.ControllerComments{
			"Home",
			`/home/:key`,
			[]string{"get"},
			nil})

}
