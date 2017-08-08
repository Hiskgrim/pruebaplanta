package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["prueba_api/controllers:DocentePlantaController"] = append(beego.GlobalControllerRouter["prueba_api/controllers:DocentePlantaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
