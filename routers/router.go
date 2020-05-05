package routers

import (
	"github.com/astaxie/beego"
	"github.com/qianbaidu/chrome-server/controllers"
)



func init() {
	//APIS
	ns :=
		beego.NewNamespace("/api/",
			beego.NSNamespace("/v1",
				//beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
				//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
				beego.NSInclude(
					&controllers.AuthController{},
					&controllers.CollectionController{},
				),
			),
		)

	beego.AddNamespace(ns)

}
