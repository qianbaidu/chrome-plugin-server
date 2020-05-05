package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"],
		beego.ControllerComments{
			Method: "DeleteUrl",
			Router: `/collection/delete-url`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"],
		beego.ControllerComments{
			Method: "EditUrl",
			Router: `/collection/edit-url/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/collection/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"],
		beego.ControllerComments{
			Method: "SaveUrl",
			Router: `/collection/save-url`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"] = append(beego.GlobalControllerRouter["github.com/qianbaidu/chrome-server/controllers:CollectionController"],
		beego.ControllerComments{
			Method: "Url",
			Router: `/collection/url/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
